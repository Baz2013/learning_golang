package lcof

import "testing"

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	target := "We%20are%20happy."

	r := ReplaceSpace(s)
	if r == target {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
