package clause

import "strings"

// Clause 是一个结构体，用于存储和管理 SQL 片段及其对应的参数。
type Clause struct {
	sql     map[Type]string        // 存储不同类型的 SQL 片段
	sqlVars map[Type][]interface{} // 存储不同类型的 SQL 参数
}

// Type 是一个自定义的整数类型，用于表示不同的 SQL 子句类型。
type Type int

// 定义各种 SQL 子句类型。
const (
	INSERT  Type = iota // INSERT 子句
	VALUES              // VALUES 子句
	SELECT              // SELECT 子句
	LIMIT               // LIMIT 子句
	WHERE               // WHERE 子句
	ORDERBY             // ORDER BY 子句
)

// Set 为 Clause 设置指定类型的 SQL 片段及其参数。
func (c *Clause) Set(name Type, vars ...interface{}) {
	// 如果 Clause 的 sql 和 sqlVars 为空，则初始化它们。
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	// 使用 generators 映射中的生成器函数生成 SQL 片段和参数。
	sql, vars := generators[name](vars...)
	// 将生成的 SQL 片段和参数存储到 Clause 中。
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 根据指定的顺序构建完整的 SQL 语句及其参数。
func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string      // 用于存储 SQL 片段
	var vars []interface{} // 用于存储 SQL 参数
	// 遍历指定的顺序，将对应的 SQL 片段和参数添加到结果中。
	for _, order := range orders {
		if sql, ok := c.sql[order]; ok { // 如果 Clause 中存在指定类型的 SQL 片段
			sqls = append(sqls, sql)                 // 将 SQL 片段添加到结果中
			vars = append(vars, c.sqlVars[order]...) // 将对应的参数添加到结果中
		}
	}
	// 将所有 SQL 片段拼接成完整的 SQL 语句，并返回 SQL 语句和参数。
	return strings.Join(sqls, " "), vars
}
