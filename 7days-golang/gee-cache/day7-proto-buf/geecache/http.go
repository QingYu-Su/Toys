package geecache

import (
	"fmt"
	"geecache/consistenthash" // 导入一致性哈希模块
	pb "geecache/geecachepb"
	"io/ioutil" // 用于读取 HTTP 响应体
	"log"       // 用于日志记录
	"net/http"  // 用于处理 HTTP 请求和响应
	"net/url"   // 用于 URL 编码
	"strings"   // 用于字符串操作
	"sync"      // 用于并发控制

	"google.golang.org/protobuf/proto"
)

const (
	defaultBasePath = "/_geecache/" // 默认的 HTTP 基础路径
	defaultReplicas = 50            // 默认的虚拟节点数量
)

// HTTPPool 是一个 HTTP 节点池，实现了 PeerPicker 接口，用于管理一组 HTTP 缓存节点。
type HTTPPool struct {
	self        string                 // 当前节点的 URL，例如 "https://example.net:8000"
	basePath    string                 // HTTP 基础路径
	mu          sync.Mutex             // 互斥锁，保护 peers 和 httpGetters
	peers       *consistenthash.Map    // 一致性哈希映射，用于选择节点
	httpGetters map[string]*httpGetter // 存储每个节点的 HTTP 客户端
}

// NewHTTPPool 初始化一个新的 HTTPPool 实例。
func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

// Log 用于记录带有服务器名称的日志。
func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

// ServeHTTP 处理所有 HTTP 请求，用于从缓存中获取数据。
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 检查请求路径是否以指定的基础路径开头
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}

	// 记录请求信息
	p.Log("%s %s", r.Method, r.URL.Path)

	// 解析请求路径，格式为 /<basepath>/<groupname>/<key>
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// 提取缓存组名和键
	groupName := parts[0] // 缓存组名
	key := parts[1]       // 缓存键

	// 获取缓存组
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	// 从缓存中获取数据
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将缓存数据序列化为协议缓冲区格式
	body, err := proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头并返回缓存数据
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
}

// Set 更新节点池中的节点列表。
func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.peers = consistenthash.New(defaultReplicas, nil)       // 初始化一致性哈希映射
	p.peers.Add(peers...)                                    // 添加节点到一致性哈希映射
	p.httpGetters = make(map[string]*httpGetter, len(peers)) // 初始化 HTTP 客户端映射

	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath} // 为每个节点创建 HTTP 客户端
	}
}

// PickPeer 根据键选择一个合适的节点。
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if peer := p.peers.Get(key); peer != "" && peer != p.self { // 选择一个节点
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true // 返回节点的 HTTP 客户端
	}
	return nil, false
}

var _ PeerPicker = (*HTTPPool)(nil) // 确保 HTTPPool 实现了 PeerPicker 接口

// httpGetter 是一个简单的 HTTP 客户端，用于从远程节点获取缓存数据。
type httpGetter struct {
	baseURL string // 远程节点的 URL
}

// Get 方法从远程节点获取缓存数据，并将结果解析为协议缓冲区格式。
func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	// 构造请求的 URL，格式为：baseURL + 编码后的组名 + 编码后的键
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,                      // 远程节点的基础 URL
		url.QueryEscape(in.GetGroup()), // 编码缓存组名
		url.QueryEscape(in.GetKey()),   // 编码缓存键
	)

	// 发送 HTTP GET 请求
	res, err := http.Get(u)
	if err != nil {
		return err // 如果请求失败，返回错误
	}
	defer res.Body.Close() // 确保在函数退出时关闭响应体

	// 检查 HTTP 响应状态码
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", res.Status) // 如果状态码不是 200，返回错误
	}

	// 读取响应体
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err) // 如果读取响应体失败，返回错误
	}

	// 将响应体解析为协议缓冲区格式
	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err) // 如果解析失败，返回错误
	}

	return nil // 如果一切正常，返回 nil
}

var _ PeerGetter = (*httpGetter)(nil) // 确保 httpGetter 实现了 PeerGetter 接口
