package session

import (
	"fmt"           // 提供格式化输入输出功能
	"geeorm/log"    // 引入日志模块，用于记录错误信息
	"geeorm/schema" // 引入 schema 模块，用于解析结构体
	"reflect"       // 提供反射功能，用于操作类型和值
	"strings"       // 提供字符串操作功能
)

// Model 设置 Session 的引用表（refTable），并返回 Session 对象。
// 如果传入的模型与当前 refTable 不一致，则重新解析模型。
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		// 如果 refTable 为空或传入的模型与当前 refTable 的模型不一致，
		// 则重新解析模型并更新 refTable。
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable 返回当前 Session 的引用表（refTable）。
// 如果 refTable 为空，则记录错误日志。
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set") // 如果 refTable 为空，记录错误日志
	}
	return s.refTable
}

// CreateTable 创建表。
// 根据 refTable 的字段信息生成 CREATE TABLE 语句并执行。
func (s *Session) CreateTable() error {
	table := s.RefTable() // 获取当前引用表
	var columns []string  // 用于存储表的列定义
	for _, field := range table.Fields {
		// 根据字段信息生成列定义，格式为 "字段名 数据类型 标签"
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",") // 将列定义拼接为字符串
	// 构造 CREATE TABLE 语句并执行
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

// DropTable 删除表。
// 根据 refTable 的表名生成 DROP TABLE 语句并执行。
func (s *Session) DropTable() error {
	// 构造 DROP TABLE 语句并执行
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

// HasTable 检查表是否存在。
// 使用方言提供的 TableExistSQL 方法生成查询语句，并检查表是否存在。
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name) // 获取检查表存在的 SQL 语句和参数
	row := s.Raw(sql, values...).QueryRow()                   // 执行查询
	var tmp string                                            // 用于存储查询结果
	_ = row.Scan(&tmp)                                        // 扫描查询结果
	return tmp == s.RefTable().Name                           // 如果查询结果等于表名，则表存在
}
