package session

import (
	"errors"
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

// Update 更新记录，支持两种参数形式：
// 1. map[string]interface{}：字段名到值的映射。
// 2. key-value 对：成对的字段名和值。
func (s *Session) Update(kv ...interface{}) (int64, error) {
	// 检查是否传入了 map[string]interface{} 类型的参数
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		// 如果不是 map 类型，则将成对的 key-value 参数转换为 map
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	// 设置 UPDATE 子句，包含表名和更新的字段映射
	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	// 构建 SQL 语句及其参数
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	// 执行 SQL 更新语句
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err // 如果执行失败，返回错误
	}
	// 返回受影响的行数
	return result.RowsAffected()
}

// Delete 删除记录。
func (s *Session) Delete() (int64, error) {
	// 设置 DELETE 子句，包含表名
	s.clause.Set(clause.DELETE, s.RefTable().Name)
	// 构建 SQL 语句及其参数
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	// 执行 SQL 删除语句
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err // 如果执行失败，返回错误
	}
	// 返回受影响的行数
	return result.RowsAffected()
}

// Count 查询记录数量。
func (s *Session) Count() (int64, error) {
	// 设置 COUNT 子句，包含表名
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	// 构建 SQL 查询语句及其参数
	sql, vars := s.clause.Build(clause.COUNT, clause.WHERE)
	// 执行查询并获取结果
	row := s.Raw(sql, vars...).QueryRow()
	var tmp int64
	if err := row.Scan(&tmp); err != nil {
		return 0, err // 如果查询失败，返回错误
	}
	// 返回查询到的数量
	return tmp, nil
}

// Limit 设置 LIMIT 子句。
func (s *Session) Limit(num int) *Session {
	// 设置 LIMIT 子句的值
	s.clause.Set(clause.LIMIT, num)
	return s // 返回 Session 实例，支持链式调用
}

// Where 添加 WHERE 子句。
func (s *Session) Where(desc string, args ...interface{}) *Session {
	// 构建 WHERE 子句的参数列表
	var vars []interface{}
	s.clause.Set(clause.WHERE, append(append(vars, desc), args...)...)
	return s // 返回 Session 实例，支持链式调用
}

// OrderBy 添加 ORDER BY 子句。
func (s *Session) OrderBy(desc string) *Session {
	// 设置 ORDER BY 子句的值
	s.clause.Set(clause.ORDERBY, desc)
	return s // 返回 Session 实例，支持链式调用
}

// First 查询第一条记录并填充到目标结构体中。
func (s *Session) First(value interface{}) error {
	// 获取目标结构体的反射值
	dest := reflect.Indirect(reflect.ValueOf(value))
	// 创建一个目标结构体类型的切片
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()
	// 使用 Limit(1) 限制查询结果为一条记录，并执行查询
	if err := s.Limit(1).Find(destSlice.Addr().Interface()); err != nil {
		return err // 如果查询失败，返回错误
	}
	// 检查查询结果是否为空
	if destSlice.Len() == 0 {
		return errors.New("NOT FOUND") // 如果未找到记录，返回错误
	}
	// 将查询结果的第一条记录赋值给目标结构体
	dest.Set(destSlice.Index(0))
	return nil
}
