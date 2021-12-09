package week2

import (
	"errors"
	"fmt"
	"testing"
)

/////////////////////////////////////////////////////////   Error vs Exception   /////////////////////////////////////////////////////////

//经常使用 errors.New() 返回一个 error 对象 (内部 errorString 对象的指针)
//如果两人创建的 error 字符内容一样，在一些场景中判定两个对象是否相等时，结果是相等的，所以每次返回一个新的对象，取其地址
//建议：errors.New(“包名: 错误信息”)

// Create a named type for our new error type
type errorString string

//Implement the error interface
func (e errorString) Error() string {
	return string(e)
}

//New creates interface values of type value
func New(text string) error {
	return errorString(text)
}

var ErrNameType = New("EOF")
var ErrStructType = errors.New("EOF")

func TestErr(t *testing.T) {
	if ErrNameType == New("EOF") {
		t.Log("Named Type Error")
	}
	if ErrStructType == errors.New("EOF") {
		t.Log("Struct Type Error")
	}

}

// C 单返回值，一般通过传递指针作为参数，返回值为 int 表示成功还是失败
// C++ 引入exception 无法知道被调用方会抛出什么异常
// Java 引入 checked Exception，方法的所有者必须申明，调用者必须处理。它们从良性到灾难性都有使用，异常的严重性由调用者来区分
// Go 多返回值和panic，让程序员知道什么时候出了问题，并为正真的异常保留了panic

func Positive(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}
	return n > -1, nil
}

func Check(n int) {
	pos, err := Positive(n)
	if err != nil {
		fmt.Println(n, err)
		return
	}
	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func TestPositive(t *testing.T) {
	Check(0)
	Check(-1)
	Check(1)
}

// 不可恢复的程序错误（索引越界、不可恢复的环境的问题）使用panic
// 其它问题使用 error
// 总结: 1 简单
//      2 考虑失败 而不是成功
//      3 没有隐藏的控制流（try catch）
//      4 完全由自己控制
//      5 Error are values
