package registry

import (
	"log"      // 提供日志记录功能
	"net/http" // 提供 HTTP 服务器和客户端功能
	"sort"     // 提供排序功能
	"strings"  // 提供字符串操作功能
	"sync"     // 提供同步原语，用于并发控制
	"time"     // 提供时间相关功能
)

// GeeRegistry 是一个简单的注册中心，提供以下功能：
// - 添加服务器并接收心跳以保持其活跃状态。
// - 同时返回所有活跃的服务器，并删除已死亡的服务器。
type GeeRegistry struct {
	timeout time.Duration          // 设置服务器的超时时间
	mu      sync.Mutex             // 保护以下字段的互斥锁
	servers map[string]*ServerItem // 存储服务器地址和其状态
}

// ServerItem 表示一个服务器的状态。
type ServerItem struct {
	Addr  string    // 服务器地址
	start time.Time // 服务器注册或心跳的时间
}

const (
	// 默认的注册中心路径
	defaultPath = "/_geerpc_/registry"
	// 默认的服务器超时时间
	defaultTimeout = time.Minute * 5
)

// New 创建一个带有超时设置的注册中心实例。
func New(timeout time.Duration) *GeeRegistry {
	return &GeeRegistry{
		servers: make(map[string]*ServerItem), // 初始化服务器映射
		timeout: timeout,                      // 设置超时时间
	}
}

// DefaultGeeRegister 是一个默认的注册中心实例，使用默认的超时时间。
var DefaultGeeRegister = New(defaultTimeout)

// putServer 将一个服务器地址添加到注册中心，或更新其心跳时间以保持活跃。
func (r *GeeRegistry) putServer(addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.servers[addr]
	if s == nil {
		// 如果服务器不存在，则添加它
		r.servers[addr] = &ServerItem{Addr: addr, start: time.Now()}
	} else {
		// 如果服务器已存在，则更新其心跳时间以保持活跃
		s.start = time.Now()
	}
}

// aliveServers 返回所有活跃的服务器地址。
// 如果设置了超时时间，则删除已死亡的服务器。
func (r *GeeRegistry) aliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	for addr, s := range r.servers {
		// 检查服务器是否在超时时间内
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, addr)
		} else {
			// 如果服务器已死亡，则从注册中心删除
			delete(r.servers, addr)
		}
	}
	// 对活跃服务器地址进行排序
	sort.Strings(alive)
	return alive
}

// ServeHTTP 实现了 http.Handler 接口，用于处理注册中心的 HTTP 请求。
// 它支持两种方法：
// - GET：返回所有活跃的服务器地址。
// - POST：添加一个服务器地址到注册中心。
func (r *GeeRegistry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// 返回所有活跃的服务器地址，通过 X-Geerpc-Servers 头部返回
		w.Header().Set("X-Geerpc-Servers", strings.Join(r.aliveServers(), ","))
	case "POST":
		// 添加一个服务器地址到注册中心
		addr := req.Header.Get("X-Geerpc-Server") // 从请求头部获取服务器地址
		if addr == "" {
			// 如果请求头部中没有 X-Geerpc-Server，返回 500 内部服务器错误
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.putServer(addr) // 将服务器地址添加到注册中心
	default:
		// 不支持的 HTTP 方法，返回 405 Method Not Allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleHTTP 注册一个 HTTP 处理器，用于处理注册中心的消息。
// 它将指定的路径映射到 GeeRegistry 的 ServeHTTP 方法。
func (r *GeeRegistry) HandleHTTP(registryPath string) {
	http.Handle(registryPath, r)                    // 将路径映射到 ServeHTTP 方法
	log.Println("rpc registry path:", registryPath) // 记录注册中心路径
}

// HandleHTTP 是 DefaultGeeRegister 的 HandleHTTP 方法的简化调用。
// 它使用默认的注册中心实例和默认路径。
func HandleHTTP() {
	DefaultGeeRegister.HandleHTTP(defaultPath) // 调用默认注册中心的 HandleHTTP 方法
}

// Heartbeat 定期发送心跳消息，以保持服务器在注册中心的活跃状态。
// 它是一个辅助函数，用于服务器注册或发送心跳。
func Heartbeat(registry, addr string, duration time.Duration) {
	if duration == 0 {
		// 如果未指定心跳间隔，则使用默认超时时间减去 1 分钟
		// 确保在服务器被删除之前有足够的时间发送心跳
		duration = defaultTimeout - time.Duration(1)*time.Minute
	}
	var err error
	err = sendHeartbeat(registry, addr) // 首次发送心跳
	go func() {
		t := time.NewTicker(duration) // 创建一个定时器
		for err == nil {
			<-t.C                               // 等待定时器触发
			err = sendHeartbeat(registry, addr) // 发送心跳
		}
	}()
}

// sendHeartbeat 向注册中心发送心跳消息。
// 它通过 HTTP POST 请求将服务器地址发送到注册中心。
func sendHeartbeat(registry, addr string) error {
	log.Println(addr, "send heart beat to registry", registry) // 记录心跳信息
	httpClient := &http.Client{}                               // 创建一个 HTTP 客户端
	req, _ := http.NewRequest("POST", registry, nil)           // 创建一个 POST 请求
	req.Header.Set("X-Geerpc-Server", addr)                    // 设置请求头部，包含服务器地址
	if _, err := httpClient.Do(req); err != nil {              // 发送请求
		log.Println("rpc server: heart beat err:", err) // 如果发送失败，记录错误
		return err
	}
	return nil // 返回 nil 表示心跳成功
}
