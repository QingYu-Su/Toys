package geerpc

import (
	"fmt"           // 提供格式化输出功能
	"html/template" // 提供 HTML 模板功能，用于生成动态网页
	"net/http"      // 提供 HTTP 服务器和客户端功能
)

// debugText 定义了一个 HTML 模板，用于显示 RPC 服务的调试信息。
// 它包含服务名称、方法名称、参数类型、返回类型以及方法调用次数。
const debugText = `<html>
	<body>
	<title>GeeRPC Services</title>
	{{range .}}
	<hr>
	Service {{.Name}}
	<hr>
		<table>
		<th align=center>Method</th><th align=center>Calls</th>
		{{range $name, $mtype := .Method}}
			<tr>
			<td align=left font=fixed>{{$name}}({{$mtype.ArgType}}, {{$mtype.ReplyType}}) error</td>
			<td align=center>{{$mtype.NumCalls}}</td>
			</tr>
		{{end}}
		</table>
	{{end}}
	</body>
	</html>`

// debug 是一个解析后的 HTML 模板，用于生成 RPC 服务的调试页面。
var debug = template.Must(template.New("RPC debug").Parse(debugText))

// debugHTTP 是一个 HTTP 处理器，用于提供 RPC 服务的调试信息。
// 它嵌入了 *Server，可以直接访问 Server 的方法和字段。
type debugHTTP struct {
	*Server
}

// debugService 是一个辅助结构体，用于存储服务的调试信息。
// 它包含服务名称和方法映射，用于在模板中渲染。
type debugService struct {
	Name   string
	Method map[string]*methodType
}

// ServeHTTP 实现了 http.Handler 接口，用于处理 /debug/geerpc 路径的请求。
// 它从 Server 的 serviceMap 中获取所有注册的服务和方法，
// 并使用模板生成一个包含调试信息的 HTML 页面。
func (server debugHTTP) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 构建一个包含所有服务调试信息的切片
	var services []debugService
	server.serviceMap.Range(func(namei, svci interface{}) bool {
		svc := svci.(*service) // 将 svci 断言为 *service 类型
		services = append(services, debugService{
			Name:   namei.(string), // 服务名称
			Method: svc.method,     // 方法映射
		})
		return true
	})
	// 使用模板生成调试页面
	err := debug.Execute(w, services)
	if err != nil {
		_, _ = fmt.Fprintln(w, "rpc: error executing template:", err.Error()) // 如果模板执行失败，记录错误
	}
}
