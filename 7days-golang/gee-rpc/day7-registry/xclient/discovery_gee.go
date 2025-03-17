package xclient

import (
	"log"      // 提供日志记录功能
	"net/http" // 提供 HTTP 客户端和服务器功能
	"strings"  // 提供字符串操作功能
	"time"     // 提供时间相关功能
)

// GeeRegistryDiscovery 是一个基于注册中心的服务发现实现。
// 它定期从注册中心刷新服务器列表，并支持随机选择和轮询选择两种模式。
type GeeRegistryDiscovery struct {
	*MultiServersDiscovery               // 嵌入 MultiServersDiscovery，复用其方法和字段
	registry               string        // 注册中心的地址
	timeout                time.Duration // 刷新服务器列表的超时时间
	lastUpdate             time.Time     // 上次更新服务器列表的时间
}

const defaultUpdateTimeout = time.Second * 10 // 默认的刷新超时时间

// Update 更新服务器列表。
func (d *GeeRegistryDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers       // 更新服务器列表
	d.lastUpdate = time.Now() // 更新最后更新时间
	return nil
}

// Refresh 从注册中心刷新服务器列表。
func (d *GeeRegistryDiscovery) Refresh() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	// 如果上次更新时间在超时时间内，则直接返回
	if d.lastUpdate.Add(d.timeout).After(time.Now()) {
		return nil
	}
	log.Println("rpc registry: refresh servers from registry", d.registry)
	resp, err := http.Get(d.registry) // 向注册中心发送 GET 请求
	if err != nil {
		log.Println("rpc registry refresh err:", err)
		return err
	}
	defer resp.Body.Close() // 确保关闭响应体
	// 从响应头部获取服务器列表
	servers := strings.Split(resp.Header.Get("X-Geerpc-Servers"), ",")
	d.servers = make([]string, 0, len(servers)) // 初始化服务器列表
	for _, server := range servers {
		if strings.TrimSpace(server) != "" { // 去除空格并过滤空字符串
			d.servers = append(d.servers, strings.TrimSpace(server))
		}
	}
	d.lastUpdate = time.Now() // 更新最后更新时间
	return nil
}

// Get 根据选择模式获取一个服务器地址。
func (d *GeeRegistryDiscovery) Get(mode SelectMode) (string, error) {
	if err := d.Refresh(); err != nil { // 刷新服务器列表
		return "", err
	}
	return d.MultiServersDiscovery.Get(mode) // 调用 MultiServersDiscovery 的 Get 方法
}

// GetAll 获取所有服务器地址。
func (d *GeeRegistryDiscovery) GetAll() ([]string, error) {
	if err := d.Refresh(); err != nil { // 刷新服务器列表
		return nil, err
	}
	return d.MultiServersDiscovery.GetAll() // 调用 MultiServersDiscovery 的 GetAll 方法
}

// NewGeeRegistryDiscovery 创建一个 GeeRegistryDiscovery 实例。
func NewGeeRegistryDiscovery(registerAddr string, timeout time.Duration) *GeeRegistryDiscovery {
	if timeout == 0 { // 如果未指定超时时间，则使用默认值
		timeout = defaultUpdateTimeout
	}
	d := &GeeRegistryDiscovery{
		MultiServersDiscovery: NewMultiServerDiscovery(make([]string, 0)), // 初始化 MultiServersDiscovery
		registry:              registerAddr,                               // 设置注册中心地址
		timeout:               timeout,                                    // 设置刷新超时时间
	}
	return d
}
