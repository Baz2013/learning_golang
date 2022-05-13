package monkey_demo

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/Baz2013/learning_golang/varys"
	"strings"
	"testing"
)

//monkey不支持内联函数，在测试的时候需要通过命令行参数-gcflags=-l关闭Go语言的内联优化。
//monkey不是线程安全的，所以不要把它用到并发的单元测试中。
//go test -run=TestMyFunc -v -gcflags=-l

func TestMyFunc(t *testing.T) {
	// 对 varys.GetInfoByUID 进行打桩
	// 无论传入的uid是多少，都返回 &varys.UserInfo{Name: "liwenzhou"}, nil
	monkey.Patch(varys.GetInfoByUID, func(int64) (*varys.UserInfo, error) {
		return &varys.UserInfo{Name: "liwenzhou"}, nil
	})

	ret := MyFunc(123)
	fmt.Println(ret)

	if !strings.Contains(ret, "liwenzhou") {
		t.Fatal()
	}
}
