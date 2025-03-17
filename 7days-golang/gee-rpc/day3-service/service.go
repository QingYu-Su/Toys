package geerpc

import (
	"go/ast"      // 提供 Go 语言的抽象语法树（AST）操作，用于检查类型和方法的可见性
	"log"         // 用于记录日志信息
	"reflect"     // 提供反射功能，用于动态操作类型和值
	"sync/atomic" // 提供原子操作，用于安全地更新和读取共享变量
)

// methodType 表示一个可调用的 RPC 方法的元数据。
type methodType struct {
	method    reflect.Method // 方法的反射表示
	ArgType   reflect.Type   // 请求参数的类型
	ReplyType reflect.Type   // 响应结果的类型
	numCalls  uint64         // 方法被调用的次数
}

// NumCalls 返回该方法被调用的次数。
func (m *methodType) NumCalls() uint64 {
	return atomic.LoadUint64(&m.numCalls) // 原子地读取调用次数
}

// newArgv 创建一个新的请求参数值。
func (m *methodType) newArgv() reflect.Value {
	var argv reflect.Value
	// 根据参数类型动态创建值：
	// 如果参数是指针类型，则创建一个指向零值的指针；
	// 如果是值类型，则直接创建零值。
	if m.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(m.ArgType.Elem())
	} else {
		argv = reflect.New(m.ArgType).Elem()
	}
	return argv
}

// newReplyv 创建一个新的响应结果值。
func (m *methodType) newReplyv() reflect.Value {
	// 响应必须是指针类型，因此直接创建一个指向零值的指针。
	// 同时，根据响应的类型初始化其值（例如，Map 或 Slice）。
	replyv := reflect.New(m.ReplyType.Elem())
	switch m.ReplyType.Elem().Kind() {
	case reflect.Map:
		// 如果响应是 Map 类型，初始化为空 Map。
		replyv.Elem().Set(reflect.MakeMap(m.ReplyType.Elem()))
	case reflect.Slice:
		// 如果响应是 Slice 类型，初始化为空 Slice。
		replyv.Elem().Set(reflect.MakeSlice(m.ReplyType.Elem(), 0, 0))
	}
	return replyv
}

// service 表示一个 RPC 服务，封装了服务的元数据和可调用的方法。
type service struct {
	name   string                 // 服务名称（通常是接收器的类型名称）
	typ    reflect.Type           // 接收器的反射类型
	rcvr   reflect.Value          // 接收器的反射值
	method map[string]*methodType // 方法名称到 methodType 的映射
}

// newService 创建一个新的 RPC 服务实例。
func newService(rcvr interface{}) *service {
	s := new(service)
	s.rcvr = reflect.ValueOf(rcvr)                  // 获取接收器的反射值
	s.name = reflect.Indirect(s.rcvr).Type().Name() // 获取服务名称
	s.typ = reflect.TypeOf(rcvr)                    // 获取接收器的反射类型
	if !ast.IsExported(s.name) {                    // 检查服务名称是否是导出的（首字母大写）
		log.Fatalf("rpc server: %s is not a valid service name", s.name)
	}
	s.registerMethods() // 注册服务中的所有方法
	return s
}

// registerMethods 遍历服务的所有方法，并注册可调用的 RPC 方法。
func (s *service) registerMethods() {
	s.method = make(map[string]*methodType)  // 初始化方法映射
	for i := 0; i < s.typ.NumMethod(); i++ { // 遍历所有方法
		method := s.typ.Method(i) // 获取方法的反射表示
		mType := method.Type      // 获取方法的类型
		// 检查方法是否符合 RPC 方法的签名：
		// - 输入参数为 3 个（接收器、请求参数、响应结果）
		// - 输出参数为 1 个（error 类型）
		if mType.NumIn() != 3 || mType.NumOut() != 1 {
			continue
		}
		if mType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() { // 检查输出是否为 error 类型
			continue
		}
		argType, replyType := mType.In(1), mType.In(2) // 获取请求参数和响应结果的类型
		// 检查请求参数和响应结果是否是导出的或内置类型
		if !isExportedOrBuiltinType(argType) || !isExportedOrBuiltinType(replyType) {
			continue
		}
		// 注册方法
		s.method[method.Name] = &methodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
		log.Printf("rpc server: register %s.%s\n", s.name, method.Name)
	}
}

// call 调用注册的 RPC 方法。
func (s *service) call(m *methodType, argv, replyv reflect.Value) error {
	atomic.AddUint64(&m.numCalls, 1)                              // 原子地增加调用次数
	f := m.method.Func                                            // 获取方法的可调用函数
	returnValues := f.Call([]reflect.Value{s.rcvr, argv, replyv}) // 调用方法
	if errInter := returnValues[0].Interface(); errInter != nil { // 检查返回的错误
		return errInter.(error)
	}
	return nil
}

// isExportedOrBuiltinType 检查类型是否是导出的或内置类型。
func isExportedOrBuiltinType(t reflect.Type) bool {
	return ast.IsExported(t.Name()) || t.PkgPath() == "" // 检查类型名称是否导出或类型是否是内置类型
}
