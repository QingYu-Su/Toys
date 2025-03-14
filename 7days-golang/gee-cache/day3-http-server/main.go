package main

import (
	"fmt"
	"geecache"
	"log"
	"net/http"
)

// 模拟一个简单的数据库，存储键值对
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	// 创建一个名为 "scores" 的缓存组，最大缓存大小为 2KB
	// 使用 geecache.GetterFunc 包装一个函数，作为数据加载器
	geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key) // 模拟从数据库加载数据
			if v, ok := db[key]; ok {
				return []byte(v), nil // 如果键存在，返回对应的值
			}
			return nil, fmt.Errorf("%s not exist", key) // 如果键不存在，返回错误
		}))

	// 定义 HTTP 服务的地址
	addr := "localhost:9999"

	// 创建一个 HTTPPool 实例，用于处理 HTTP 请求
	peers := geecache.NewHTTPPool(addr)

	// 启动 HTTP 服务，监听指定地址
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
