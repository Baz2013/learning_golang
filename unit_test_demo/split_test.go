package unit_test_demo

import (
	"reflect"
	"testing"
)

// 在目标目录下执行单元测试命令
// 命令: go test
// go test -v # 输出详细信息
// go test -run=xx -v # -run 参数可以正则匹配测试方法名
// 例如： go test -run=Sep -v # 在这里可以匹配到TestSplitWithComplexSep()方法

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能直接比较，需借助反射包中的方法比较
		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

func TestSplitWithComplexSep(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 表格驱动测试
// 官方标准库中有很多表格驱动测试的示例，例如fmt包中的测试代码
func TestSplitAll(t *testing.T) {
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}

// 使用工具生成测试代码 gotests
// 安装： go get -u github.com/cweill/gotests/...
// gotests -all -w split.go
// 上面的命令表示，为split.go文件的所有函数生成测试代码至split_test.go文件（目录下如果事先存在这个文件就不再生成）。

// 测试覆盖率
// go test -cover
// go test -cover -coverprofile=out  #用来将覆盖率相关的记录信息输出到一个文件
// go tool cover -html=out  #使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告

// testify是一个社区非常流行的Go单元测试工具包，其中使用最多的功能就是它提供的断言工具——testify/assert或testify/require。
