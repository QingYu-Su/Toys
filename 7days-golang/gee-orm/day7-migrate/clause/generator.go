package clause

import (
	"fmt"
	"strings"
)

// generator 是一个函数类型，用于生成 SQL 片段及其参数。
type generator func(values ...interface{}) (string, []interface{})

// generators 是一个映射，将 Clause 类型映射到对应的生成器函数。
var generators map[Type]generator

// init 初始化 generators 映射，注册各种 Clause 类型的生成器函数。
func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert   // 注册 INSERT 类型的生成器
	generators[VALUES] = _values   // 注册 VALUES 类型的生成器
	generators[SELECT] = _select   // 注册 SELECT 类型的生成器
	generators[LIMIT] = _limit     // 注册 LIMIT 类型的生成器
	generators[WHERE] = _where     // 注册 WHERE 类型的生成器
	generators[ORDERBY] = _orderBy // 注册 ORDER BY 类型的生成器
	generators[UPDATE] = _update   // 注册 UPDATE 类型的生成器
	generators[DELETE] = _delete   // 注册 DELETE 类型的生成器
	generators[COUNT] = _count     // 注册 COUNT 类型的生成器
}

// genBindVars 生成指定数量的占位符字符串，用于 SQL 参数绑定。
func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

// _insert 生成 INSERT INTO 语句的 SQL 片段。
func _insert(values ...interface{}) (string, []interface{}) {
	tableName := values[0]                            // 表名
	fields := strings.Join(values[1].([]string), ",") // 字段列表
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

// _values 生成 VALUES 语句的 SQL 片段。
func _values(values ...interface{}) (string, []interface{}) {
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("VALUES ")
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v)) // 生成占位符字符串
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

// _select 生成 SELECT 语句的 SQL 片段。
func _select(values ...interface{}) (string, []interface{}) {
	tableName := values[0]                            // 表名
	fields := strings.Join(values[1].([]string), ",") // 字段列表
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

// _limit 生成 LIMIT 语句的 SQL 片段。
func _limit(values ...interface{}) (string, []interface{}) {
	return "LIMIT ?", values
}

// _where 生成 WHERE 语句的 SQL 片段。
func _where(values ...interface{}) (string, []interface{}) {
	desc, vars := values[0], values[1:] // 条件描述和参数
	return fmt.Sprintf("WHERE %s", desc), vars
}

// _orderBy 生成 ORDER BY 语句的 SQL 片段。
func _orderBy(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

// _update 生成 UPDATE 语句的 SQL 片段。
func _update(values ...interface{}) (string, []interface{}) {
	tableName := values[0]                  // 表名
	m := values[1].(map[string]interface{}) // 更新的字段和值
	var keys []string
	var vars []interface{}
	for k, v := range m {
		keys = append(keys, k+" = ?") // 生成字段更新表达式
		vars = append(vars, v)        // 收集参数
	}
	return fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(keys, ", ")), vars
}

// _delete 生成 DELETE FROM 语句的 SQL 片段。
func _delete(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("DELETE FROM %s", values[0]), []interface{}{}
}

// _count 生成 COUNT(*) 查询的 SQL 片段。
func _count(values ...interface{}) (string, []interface{}) {
	return _select(values[0], []string{"count(*)"}) // 调用 _select 生成 COUNT 查询
}
