package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 默认的 HTTP 基础路径，用于访问缓存服务
const defaultBasePath = "/_geecache/"

// HTTPPool 表示一个 HTTP 缓存服务池，管理一个缓存服务的 HTTP 接口。
type HTTPPool struct {
	self     string // 当前服务的地址（如 http://localhost:8001）
	basePath string // 访问缓存服务的基础路径
}

// NewHTTPPool 创建一个新的 HTTPPool 实例。
func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath, // 默认使用全局定义的基础路径
	}
}

// Log 方法用于记录日志，方便调试和监控。
func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

// ServeHTTP 方法实现了 http.Handler 接口，用于处理 HTTP 请求。
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 检查请求路径是否以指定的基础路径开头
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}

	// 记录请求的日志
	p.Log("%s %s", r.Method, r.URL.Path)

	// 去掉基础路径后，解析请求路径为缓存组名和键
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)

	// 确保路径格式正确（必须包含缓存组名和键）
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0] // 缓存组名
	key := parts[1]       // 缓存键

	// 从全局缓存组映射中获取指定的缓存组
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	// 从缓存组中获取缓存值
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置 HTTP 响应头，并返回缓存值
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}
