package session

import (
	"geeorm/clause"
	"reflect"
)

// Insert 插入一条或多条记录到数据库中。
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0) // 用于存储所有记录的字段值
	for _, value := range values {
		// 获取当前记录对应的表结构
		table := s.Model(value).RefTable()
		// 设置 INSERT 子句，包含表名和字段名
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		// 提取当前记录的字段值，并追加到 recordValues 中
		recordValues = append(recordValues, table.RecordValues(value))
	}

	// 设置 VALUES 子句，包含所有记录的字段值
	s.clause.Set(clause.VALUES, recordValues...)
	// 构建完整的 SQL 插入语句及其参数
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	// 执行 SQL 插入语句
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err // 如果执行失败，返回错误
	}

	// 返回受影响的行数
	return result.RowsAffected()
}

// Find 查询记录并填充到目标切片中。
func (s *Session) Find(values interface{}) error {
	// 获取目标切片的反射值，并去除指针获取实际切片
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// 获取目标切片的元素类型
	destType := destSlice.Type().Elem()
	// 创建一个新的目标实例，并获取其对应的表结构
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()
	// 设置 SELECT 子句，包含表名和字段名
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	// 构建查询 SQL 语句及其参数
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	// 执行查询并获取结果集
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err // 如果查询失败，返回错误
	}

	// 遍历结果集
	for rows.Next() {
		// 创建一个新的目标实例
		dest := reflect.New(destType).Elem()
		// 提取目标实例的字段地址，用于 Scan 方法
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		// 将查询结果扫描到目标实例的字段中
		if err := rows.Scan(values...); err != nil {
			return err // 如果扫描失败，返回错误
		}
		// 将目标实例追加到结果切片中
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	// 关闭结果集
	return rows.Close()
}
