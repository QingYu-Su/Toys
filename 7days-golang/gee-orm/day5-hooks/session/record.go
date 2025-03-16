package session

import (
	"errors"
	"geeorm/clause"
	"reflect"
)

// Insert 插入一条或多条记录到数据库中。
// 支持传入多个结构体实例作为参数。
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0) // 用于存储所有记录的字段值
	for _, value := range values {
		// 在插入前调用 BeforeInsert 钩子方法
		s.CallMethod(BeforeInsert, value)
		table := s.Model(value).RefTable()                             // 获取记录对应的表结构
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)      // 设置 INSERT 子句
		recordValues = append(recordValues, table.RecordValues(value)) // 提取并存储记录的字段值
	}

	// 设置 VALUES 子句
	s.clause.Set(clause.VALUES, recordValues...)
	// 构建完整的 SQL 插入语句及其参数
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	// 执行 SQL 插入语句
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err // 如果执行失败，返回错误
	}
	// 在插入后调用 AfterInsert 钩子方法
	s.CallMethod(AfterInsert, nil)
	// 返回受影响的行数
	return result.RowsAffected()
}

// Find 查询符合条件的记录，并填充到目标切片中。
func (s *Session) Find(values interface{}) error {
	// 在查询前调用 BeforeQuery 钩子方法
	s.CallMethod(BeforeQuery, nil)
	destSlice := reflect.Indirect(reflect.ValueOf(values))                // 获取目标切片的反射值
	destType := destSlice.Type().Elem()                                   // 获取目标切片的元素类型
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable() // 获取表结构

	// 设置 SELECT 子句
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
		dest := reflect.New(destType).Elem() // 创建一个新的目标实例
		var values []interface{}
		// 提取目标实例的字段地址，用于 Scan 方法
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		// 将查询结果扫描到目标实例的字段中
		if err := rows.Scan(values...); err != nil {
			return err // 如果扫描失败，返回错误
		}
		// 在查询后调用 AfterQuery 钩子方法
		s.CallMethod(AfterQuery, dest.Addr().Interface())
		// 将目标实例追加到结果切片中
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	// 关闭结果集
	return rows.Close()
}

// First 查询第一条记录，并填充到目标结构体中。
func (s *Session) First(value interface{}) error {
	dest := reflect.Indirect(reflect.ValueOf(value))              // 获取目标结构体的反射值
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem() // 创建目标结构体类型的切片
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

// Limit 设置 LIMIT 子句，限制查询结果的数量。
func (s *Session) Limit(num int) *Session {
	// 设置 LIMIT 子句的值
	s.clause.Set(clause.LIMIT, num)
	return s // 返回 Session 实例，支持链式调用
}

// Where 添加 WHERE 子句，设置查询条件。
func (s *Session) Where(desc string, args ...interface{}) *Session {
	var vars []interface{}
	// 构建 WHERE 子句的参数列表
	s.clause.Set(clause.WHERE, append(append(vars, desc), args...)...)
	return s // 返回 Session 实例，支持链式调用
}

// OrderBy 添加 ORDER BY 子句，设置排序条件。
func (s *Session) OrderBy(desc string) *Session {
	// 设置 ORDER BY 子句的值
	s.clause.Set(clause.ORDERBY, desc)
	return s // 返回 Session 实例，支持链式调用
}

// Update 更新记录，支持两种参数形式：
// 1. map[string]interface{}：字段名到值的映射。
// 2. key-value 对：成对的字段名和值。
func (s *Session) Update(kv ...interface{}) (int64, error) {
	// 在更新前调用 BeforeUpdate 钩子方法
	s.CallMethod(BeforeUpdate, nil)
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	// 设置 UPDATE 子句
	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	// 构建 SQL 更新语句及其参数
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	// 执行 SQL 更新语句
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err // 如果执行失败，返回错误
	}
	// 在更新后调用 AfterUpdate 钩子方法
	s.CallMethod(AfterUpdate, nil)
	// 返回受影响的行数
	return result.RowsAffected()
}

// Delete 删除记录。
func (s *Session) Delete() (int64, error) {
	// 在删除前调用 BeforeDelete 钩子方法
	s.CallMethod(BeforeDelete, nil)
	// 设置 DELETE 子句
	s.clause.Set(clause.DELETE, s.RefTable().Name)
	// 构建 SQL 删除语句及其参数
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	// 执行 SQL 删除语句
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err // 如果执行失败，返回错误
	}
	// 在删除后调用 AfterDelete 钩子方法
	s.CallMethod(AfterDelete, nil)
	// 返回受影响的行数
	return result.RowsAffected()
}

// Count 查询符合条件的记录数量。
func (s *Session) Count() (int64, error) {
	// 设置 COUNT 子句
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	// 构建查询 SQL 语句及其参数
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
