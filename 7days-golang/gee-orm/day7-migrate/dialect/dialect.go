package dialect

import "reflect" // 提供反射功能，用于操作类型和值

// dialectsMap 是一个全局的映射，用于存储已注册的方言（Dialect）。
var dialectsMap = map[string]Dialect{}

// Dialect 定义了数据库方言的接口，用于封装不同数据库的特定行为。
type Dialect interface {
	// DataTypeOf 返回给定 Go 类型对应的数据库数据类型。
	DataTypeOf(typ reflect.Value) string

	// TableExistSQL 返回用于检查表是否存在的 SQL 查询语句及其参数。
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect 用于注册一个新的方言到全局映射中。
// 参数 name 是方言的名称，dialect 是实现了 Dialect 接口的具体方言实例。
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect // 将方言实例存储到映射中
}

// GetDialect 用于从全局映射中获取已注册的方言。
// 参数 name 是方言的名称，返回值 dialect 是对应的方言实例，ok 表示是否找到。
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name] // 从映射中查找方言
	return
}
