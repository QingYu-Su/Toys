package dialect

import (
	"fmt"     // 提供格式化输入输出功能
	"reflect" // 提供反射功能，用于操作类型和值
	"time"    // 提供时间相关的功能
)

// sqlite3 是一个实现了 Dialect 接口的结构体，用于封装 SQLite3 数据库的特定行为。
type sqlite3 struct{}

// 确保 sqlite3 实现了 Dialect 接口。
var _ Dialect = (*sqlite3)(nil)

// 在包初始化时注册 sqlite3 方言。
func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

// DataTypeOf 根据反射值返回 SQLite3 数据库中对应的 SQL 数据类型。
func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() { // 根据反射值的类型种类返回对应的 SQL 数据类型
	case reflect.Bool:
		return "bool" // 对应 SQLite3 的布尔类型
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer" // 对应 SQLite3 的整数类型
	case reflect.Int64, reflect.Uint64:
		return "bigint" // 对应 SQLite3 的大整数类型
	case reflect.Float32, reflect.Float64:
		return "real" // 对应 SQLite3 的浮点数类型
	case reflect.String:
		return "text" // 对应 SQLite3 的文本类型
	case reflect.Array, reflect.Slice:
		return "blob" // 对应 SQLite3 的二进制数据类型
	case reflect.Struct:
		// 特殊处理 time.Time 类型
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime" // 对应 SQLite3 的日期时间类型
		}
	}
	// 如果无法识别的类型，抛出 panic
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

// TableExistSQL 返回用于检查 SQLite3 数据库中表是否存在的 SQL 查询语句及其参数。
func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName} // SQL 查询参数
	// 构造用于检查表是否存在的 SQL 查询语句
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
