package leetcode_cn

import "testing"

func TestFirstBadVersion(t *testing.T) {
	n := firstBadVersion(5)
	if n == 4 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
