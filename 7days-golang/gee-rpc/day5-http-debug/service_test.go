package geerpc

import (
	"fmt"     // 用于格式化输出和错误信息
	"reflect" // 提供反射功能，用于动态操作类型和值
	"testing" // 提供测试框架功能
)

// Foo 是一个用于测试的结构体，表示一个 RPC 服务的接收器。
// 它包含两个方法：`Sum` 和 `sum`，分别用于测试导出方法和未导出方法的行为。
type Foo int

// Args 是一个用于测试的结构体，表示 RPC 方法的请求参数。
type Args struct{ Num1, Num2 int }

// Sum 是一个导出的方法，符合 RPC 方法的签名。
// 它接收一个 `Args` 类型的参数，计算两个数字的和，并将结果存储在 `reply` 指针指向的变量中。
func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2 // 计算结果并存储到 reply 中
	return nil                     // 返回 nil 表示调用成功
}

// sum 是一个未导出的方法，用于测试未导出方法是否会被注册为 RPC 方法。
// 它的行为与 `Sum` 方法相同，但由于未导出，不应被注册为 RPC 方法。
func (f Foo) sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

// _assert 是一个辅助函数，用于断言条件是否成立。
// 如果条件不成立，它会抛出一个 panic，包含错误信息和可选的格式化参数。
func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v...))
	}
}

// TestNewService 测试 newService 函数是否正确注册了服务的方法。
func TestNewService(t *testing.T) {
	var foo Foo                                                                              // 创建一个 Foo 类型的变量
	s := newService(&foo)                                                                    // 使用 Foo 的指针创建一个服务实例
	_assert(len(s.method) == 1, "wrong service Method, expect 1, but got %d", len(s.method)) // 断言服务中只有一个方法
	mType := s.method["Sum"]                                                                 // 获取名为 "Sum" 的方法类型
	_assert(mType != nil, "wrong Method, Sum shouldn't nil")                                 // 断言 "Sum" 方法不为 nil
}

// TestMethodType_Call 测试 methodType 的 Call 方法是否正确调用了 RPC 方法。
func TestMethodType_Call(t *testing.T) {
	var foo Foo              // 创建一个 Foo 类型的变量
	s := newService(&foo)    // 使用 Foo 的指针创建一个服务实例
	mType := s.method["Sum"] // 获取名为 "Sum" 的方法类型

	// 创建请求参数和响应结果的反射值
	argv := mType.newArgv()
	replyv := mType.newReplyv()

	// 设置请求参数的值
	argv.Set(reflect.ValueOf(Args{Num1: 1, Num2: 3}))

	// 调用方法并检查结果
	err := s.call(mType, argv, replyv)
	_assert(err == nil && *replyv.Interface().(*int) == 4 && mType.NumCalls() == 1, "failed to call Foo.Sum") // 断言调用成功且结果正确
}
