package main

import (
	"flag" // 用于解析命令行参数
	"fmt"
	"geecache" // 自定义的缓存模块
	"log"
	"net/http"
)

// 模拟一个简单的数据库，存储键值对
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// createGroup 创建一个缓存组，使用自定义的 GetterFunc 作为数据加载器。
func createGroup() *geecache.Group {
	return geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key) // 模拟从数据库加载数据
			if v, ok := db[key]; ok {
				return []byte(v), nil // 如果键存在，返回对应的值
			}
			return nil, fmt.Errorf("%s not exist", key) // 如果键不存在，返回错误
		}))
}

// startCacheServer 启动一个缓存服务，监听指定的地址，并注册到一致性哈希环中。
func startCacheServer(addr string, addrs []string, gee *geecache.Group) {
	peers := geecache.NewHTTPPool(addr) // 创建一个 HTTP 节点池
	peers.Set(addrs...)                 // 设置节点池中的节点列表
	gee.RegisterPeers(peers)            // 将节点池注册到缓存组中
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers)) // 启动 HTTP 服务
}

// startAPIServer 启动一个 API 服务，提供缓存数据的 HTTP 接口。
func startAPIServer(apiAddr string, gee *geecache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key") // 从请求中获取键
			view, err := gee.Get(key)       // 从缓存组中获取数据
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError) // 如果出错，返回错误
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream") // 设置响应头
			w.Write(view.ByteSlice())                                  // 返回缓存数据
		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil)) // 启动 HTTP 服务
}

func main() {
	var port int                                             // 定义缓存服务的端口
	var api bool                                             // 是否启动 API 服务
	flag.IntVar(&port, "port", 8001, "Geecache server port") // 默认端口为 8001
	flag.BoolVar(&api, "api", false, "Start a api server?")  // 默认不启动 API 服务
	flag.Parse()                                             // 解析命令行参数

	apiAddr := "http://localhost:9999" // API 服务的地址
	addrMap := map[int]string{         // 缓存服务的地址映射
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string // 存储所有缓存服务的地址
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup() // 创建缓存组
	if api {
		go startAPIServer(apiAddr, gee) // 启动 API 服务
	}
	startCacheServer(addrMap[port], addrs, gee) // 启动缓存服务
}
