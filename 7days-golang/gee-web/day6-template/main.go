package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"gee" // 导入自定义的 gee 模块（需修改 go.mod 以确保正确导入本地模块）
)

// student 定义了一个学生结构体
type student struct {
	Name string
	Age  int8
}

// FormatAsDate 定义了一个自定义模板函数，用于格式化时间
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New() // 创建一个新的 gee.Engine 实例

	// 使用全局中间件 Logger
	r.Use(gee.Logger())

	// 设置自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})

	// 加载 HTML 模板文件
	r.LoadHTMLGlob("templates/*")

	// 提供静态文件服务
	r.Static("/assets", "./static")

	// 创建学生实例
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	// 定义路由，渲染 CSS 模板
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	// 定义路由，渲染数组模板
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	// 定义路由，渲染自定义函数模板
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	// 启动服务器
	r.Run(":9999")
}
