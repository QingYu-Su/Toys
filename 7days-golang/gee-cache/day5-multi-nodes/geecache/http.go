package geecache

import (
	"fmt"
	"geecache/consistenthash" // 导入一致性哈希模块
	"io/ioutil"               // 用于读取 HTTP 响应体
	"log"                     // 用于日志记录
	"net/http"                // 用于处理 HTTP 请求和响应
	"net/url"                 // 用于 URL 编码
	"strings"                 // 用于字符串操作
	"sync"                    // 用于并发控制
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
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}
	p.Log("%s %s", r.Method, r.URL.Path)

	// 解析请求路径，格式为 /<basepath>/<groupname>/<key>
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0] // 缓存组名
	key := parts[1]       // 缓存键

	group := GetGroup(groupName) // 获取缓存组
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key) // 从缓存中获取数据
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头并返回缓存数据
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
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

// Get 从远程节点获取缓存数据。
func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(group), // 编码缓存组名
		url.QueryEscape(key),   // 编码缓存键
	)
	res, err := http.Get(u) // 发送 HTTP 请求
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body) // 读取响应体
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	return bytes, nil
}

var _ PeerGetter = (*httpGetter)(nil) // 确保 httpGetter 实现了 PeerGetter 接口
