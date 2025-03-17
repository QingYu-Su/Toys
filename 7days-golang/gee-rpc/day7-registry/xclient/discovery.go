package xclient

import (
	"errors"    // 提供错误处理功能
	"math"      // 提供数学函数，例如随机数生成
	"math/rand" // 提供随机数生成器
	"sync"      // 提供同步原语，用于并发控制
	"time"      // 提供时间相关功能，用于生成随机种子
)

// SelectMode 定义了选择服务器的模式。
type SelectMode int

const (
	// RandomSelect 表示随机选择服务器
	RandomSelect SelectMode = iota
	// RoundRobinSelect 表示使用轮询算法选择服务器
	RoundRobinSelect
)

// Discovery 定义了一个服务发现接口，用于动态获取和更新服务器列表。
type Discovery interface {
	// Refresh 从远程注册中心刷新服务器列表
	Refresh() error
	// Update 动态更新服务器列表
	Update(servers []string) error
	// Get 根据选择模式获取一个服务器地址
	Get(mode SelectMode) (string, error)
	// GetAll 获取所有服务器地址
	GetAll() ([]string, error)
}

// MultiServersDiscovery 是一个不依赖注册中心的多服务器发现实现。
// 用户需要显式提供服务器地址。
type MultiServersDiscovery struct {
	r       *rand.Rand   // 用于生成随机数
	mu      sync.RWMutex // 保护以下字段
	servers []string     // 服务器地址列表
	index   int          // 轮询算法中记录的当前选择位置
}

var _ Discovery = (*MultiServersDiscovery)(nil)

// Refresh 对于 MultiServersDiscovery 没有意义，因此直接返回 nil。
func (d *MultiServersDiscovery) Refresh() error {
	return nil
}

// Update 动态更新服务器列表。
func (d *MultiServersDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers // 更新服务器列表
	return nil
}

// Get 根据选择模式获取一个服务器地址。
func (d *MultiServersDiscovery) Get(mode SelectMode) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers") // 如果没有服务器可用，返回错误
	}
	switch mode {
	case RandomSelect:
		return d.servers[d.r.Intn(n)], nil // 随机选择一个服务器
	case RoundRobinSelect:
		s := d.servers[d.index%n]   // 使用轮询算法选择服务器
		d.index = (d.index + 1) % n // 更新索引
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode") // 不支持的选择模式
	}
}

// GetAll 返回发现中的所有服务器地址。
func (d *MultiServersDiscovery) GetAll() ([]string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	// 返回服务器列表的副本
	servers := make([]string, len(d.servers))
	copy(servers, d.servers)
	return servers, nil
}

// NewMultiServerDiscovery 创建一个 MultiServersDiscovery 实例。
func NewMultiServerDiscovery(servers []string) *MultiServersDiscovery {
	d := &MultiServersDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())), // 使用当前时间戳作为随机种子
	}
	d.index = d.r.Intn(math.MaxInt32 - 1) // 初始化轮询索引
	return d
}
