package week2

import (
	"fmt"
	"testing"
)

/////////////////////////////////////////////////////////   Error Type   /////////////////////////////////////////////////////////

// Sentinel Error (预定义的特定错误)
// 不灵活
// Error方法服务于程序员而不是程序，业务代码不能依赖其输出
// 成为API公共部分：会增加接口的表面积
// 在两个包之间创建了依赖
// 尽量避免使用sentinel Error

// Error types
// 实现了error接口的自定义类型
// 优势：能够包装底层错误以提供更多上下层
// 缺点：调用者使用类型断言和类型switch，就要让自定义的error变为public，导致和调用者产生强耦合，导致API变得脆弱。
// 避免使用

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func Test1(t *testing.T) {
	t.Error(&MyError{"Something happened", "server.go", 88})
}

// Opaque errors (不透明的错误处理)
// 最灵活：只返回错误，不假设内容
// 少数情况下，二分错误处理方法是不够的。我们可以断言错误实现了特定的行为，而不是断言错误是特定的类型或值。

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}
