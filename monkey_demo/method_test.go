package monkey_demo

import (
	"bou.ke/monkey"
	"reflect"
	"strings"
	"testing"
)

//如果我们为GetInfo编写单元测试的时候CalcAge方法的功能还未完成，这个时候我们可以使用monkey进行打桩。

func TestUser_GetInfo(t *testing.T) {
	var u = &User{
		Name:     "q1mi",
		Birthday: "1990-12-20",
	}

	// 为对象方法打桩
	monkey.PatchInstanceMethod(reflect.TypeOf(u), "CalcAge", func(*User) int {
		return 18
	})

	ret := u.GetInfo() // 内部调用u.CalcAge方法时会返回18
	if !strings.Contains(ret, "朋友") {
		t.Fatal()
	}
}
