package gee

import (
	"fmt"
	"reflect"
	"testing"
)

// newTestRouter 创建一个测试用的路由实例，包含几个预定义的路由
func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)                 // 添加根路径路由
	r.addRoute("GET", "/hello/:name", nil)      // 添加带有动态参数的路由
	r.addRoute("GET", "/hello/b/c", nil)        // 添加固定路径路由
	r.addRoute("GET", "/hi/:name", nil)         // 添加另一个带有动态参数的路由
	r.addRoute("GET", "/assets/*filepath", nil) // 添加带有通配符参数的路由
	return r
}

// TestParsePattern 测试 parsePattern 函数是否正确拆分路径模式
func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})        // 测试动态参数
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})           // 测试单个通配符
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"}) // 测试多个通配符

	if !ok {
		t.Fatal("test parsePattern failed") // 如果测试失败，报告错误
	}
}

// TestGetRoute 测试 getRoute 方法是否正确匹配路径并提取参数
func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu") // 测试匹配带有动态参数的路径

	if n == nil {
		t.Fatal("nil shouldn't be returned") // 如果返回 nil，报告错误
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name") // 如果匹配的路径模式不正确，报告错误
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'") // 如果提取的参数不正确，报告错误
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"]) // 打印匹配结果
}

// TestGetRoute2 测试 getRoute 方法是否正确匹配带有通配符参数的路径
func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file1.txt") // 测试匹配带有通配符参数的路径
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("pattern should be /assets/*filepath & filepath should be file1.txt") // 如果匹配结果不正确，报告错误
	}

	n2, ps2 := r.getRoute("GET", "/assets/css/test.css") // 测试匹配带有通配符参数的路径
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern should be /assets/*filepath & filepath should be css/test.css") // 如果匹配结果不正确，报告错误
	}
}

// TestGetRoutes 测试 getRoutes 方法是否正确返回所有路由
func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET") // 获取所有 GET 路由
	for i, n := range nodes {
		fmt.Println(i+1, n) // 打印每个路由节点
	}

	if len(nodes) != 5 {
		t.Fatal("the number of routes should be 5") // 如果路由数量不正确，报告错误
	}
}
