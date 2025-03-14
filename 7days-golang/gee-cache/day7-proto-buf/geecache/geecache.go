package geecache

import (
	"fmt"
	pb "geecache/geecachepb"
	"geecache/singleflight"
	"log"
	"sync"
)

// Getter 是一个接口，定义了从外部数据源加载数据的方法。
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc 是一个函数类型，实现了 Getter 接口。
type GetterFunc func(key string) ([]byte, error)

// Get 方法使得 GetterFunc 类型的函数可以满足 Getter 接口。
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// Group 是一个缓存组，管理一个缓存实例和一个外部数据加载器。
type Group struct {
	name      string              // 缓存组的名称，用于标识不同的缓存组
	getter    Getter              // 外部数据加载器，用于从外部数据源加载数据
	mainCache cache               // 主缓存实例，用于存储本地缓存数据
	peers     PeerPicker          // 节点选择器，用于分布式缓存，选择其他节点获取数据
	loader    *singleflight.Group // 用于避免重复加载相同的数据，确保并发请求时只执行一次加载逻辑
}

// 全局变量，用于存储所有缓存组
var (
	mu     sync.RWMutex              // 读写互斥锁，保护对 groups 的并发访问
	groups = make(map[string]*Group) // 存储所有缓存组的映射
)

// NewGroup 创建一个新的缓存组，并将其注册到全局 map 中。
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter") // 如果 getter 为空，抛出 panic
	}
	mu.Lock()         // 加写锁
	defer mu.Unlock() // 确保在函数退出时释放锁

	g := &Group{
		name:      name,                          // 设置缓存组的名称
		getter:    getter,                        // 设置外部数据加载器
		mainCache: cache{cacheBytes: cacheBytes}, // 初始化主缓存实例，并设置缓存大小
		loader:    &singleflight.Group{},         // 初始化 singleflight.Group，用于避免重复加载
	}
	groups[name] = g // 将缓存组注册到全局 map 中
	return g
}

// GetGroup 从全局 map 中获取一个缓存组。
func GetGroup(name string) *Group {
	mu.RLock() // 加读锁
	g := groups[name]
	mu.RUnlock() // 确保在函数退出时释放锁
	return g
}

// Get 方法尝试从缓存中获取一个键的值。
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required") // 如果键为空，返回错误
	}

	if v, ok := g.mainCache.get(key); ok { // 尝试从缓存中获取
		log.Println("[GeeCache] hit") // 日志记录：缓存命中
		return v, nil
	}

	return g.load(key) // 缓存未命中，从外部加载
}

// load 方法从外部数据源加载数据，并将其存储到缓存中。
func (g *Group) load(key string) (value ByteView, err error) {
	// 使用 singleflight.Group 来避免重复加载相同的数据
	viewi, err := g.loader.Do(key, func() (interface{}, error) {
		// 如果有分布式缓存节点，尝试从其他节点获取数据
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok { // 选择一个节点
				if value, err = g.getFromPeer(peer, key); err == nil { // 从该节点获取数据
					return value, nil
				}
				log.Println("[GeeCache] Failed to get from peer", err) // 如果失败，记录日志
			}
		}

		// 如果没有分布式节点，或者从分布式节点获取失败，则从本地加载
		return g.getLocally(key)
	})

	// 如果加载成功，将结果转换为 ByteView 类型并返回
	if err == nil {
		return viewi.(ByteView), nil
	}
	return
}

// getFromPeer 从远程节点获取数据。
func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	req := &pb.Request{
		Group: g.name,
		Key:   key,
	}

	res := &pb.Response{}
	err := peer.Get(req, res)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: res.Value}, nil
}

// getLocally 方法从外部数据源加载数据。
func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key) // 调用 Getter 接口加载数据
	if err != nil {
		return ByteView{}, err // 如果加载失败，返回错误
	}

	value := ByteView{b: cloneBytes(bytes)} // 创建 ByteView 实例
	g.populateCache(key, value)             // 将数据存储到缓存中
	return value, nil
}

// populateCache 方法将数据存储到缓存中。
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value) // 调用缓存实例的 add 方法
}

// RegisterPeers 注册一个节点选择器，用于分布式缓存。
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once") // 防止重复注册
	}
	g.peers = peers
}
