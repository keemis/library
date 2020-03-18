package errs

import (
	"testing"
)

func TestUse(t *testing.T) {
	// 普通
	err := New("an exception occurred")
	t.Logf("err: %v ", err.Error())
	t.Logf("err: %v ", err)
	// 普通 + 堆栈
	t.Logf("err: %+v ", err)

	// Errorf 用法
	err = Errorf("an exception occurred: %v", "php language")
	t.Logf("err: %v ", err.Error())

	// Wrap 用法
	err = Wrap(err, "problem")
	t.Logf("err: %v ", err.Error())

	// 高级用法
	err = NewWithOption(WithCode(501), WithMsg("%v high-level exceptions", "two"), WithData("go go"))
	t.Logf("err: %v ", err)
	t.Logf("code: %v , msg: %v, data: %v", GetCode(err), GetMsg(err), GetData(err))
	t.Logf("stack: %v", GetStack(err))
}
