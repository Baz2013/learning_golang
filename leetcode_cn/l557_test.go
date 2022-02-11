package leetcode_cn

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	expectedResult := "s'teL ekat edoCteeL tsetnoc"
	result := reverseWords(s)
	fmt.Println(result)
	if result == expectedResult {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	s = "God Ding"
	expectedResult = "doG gniD"
	result = reverseWords(s)
	if result == expectedResult {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
