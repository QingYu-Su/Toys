package session

import (
	"geeorm/log"
	"reflect"
)

// Hooks constants 定义了支持的钩子方法名称。
const (
	BeforeQuery  = "BeforeQuery"  // 查询前调用
	AfterQuery   = "AfterQuery"   // 查询后调用
	BeforeUpdate = "BeforeUpdate" // 更新前调用
	AfterUpdate  = "AfterUpdate"  // 更新后调用
	BeforeDelete = "BeforeDelete" // 删除前调用
	AfterDelete  = "AfterDelete"  // 删除后调用
	BeforeInsert = "BeforeInsert" // 插入前调用
	AfterInsert  = "AfterInsert"  // 插入后调用
)

// CallMethod 调用指定的钩子方法。
// 参数：
//   - method：钩子方法的名称（如 BeforeQuery、AfterInsert 等）。
//   - value：可选的目标对象，用于调用其上的钩子方法。如果为 nil，则调用 RefTable().Model 上的方法。
func (s *Session) CallMethod(method string, value interface{}) {
	// 如果 value 不为 nil，则从 value 中获取方法；否则从 RefTable().Model 中获取方法。
	var fm reflect.Value
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method) // 从 value 中获取方法
	} else {
		fm = reflect.ValueOf(s.RefTable().Model).MethodByName(method) // 从 RefTable().Model 中获取方法
	}

	// 构造方法调用的参数（当前 Session 实例）
	param := []reflect.Value{reflect.ValueOf(s)}

	// 检查方法是否存在
	if fm.IsValid() {
		// 调用方法，并获取返回值
		if v := fm.Call(param); len(v) > 0 {
			// 如果方法返回值的第一个参数是 error 类型，则记录错误
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
}
