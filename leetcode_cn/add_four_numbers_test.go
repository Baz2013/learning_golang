package leetcode_cn

import "testing"

func TestAddFourNumbers(t *testing.T) {
	A := [3]int{3, 1, 4}
	B := [3]int{-2, 4, 3}
	C := [3]int{-1, 3, 2}
	D := [3]int{3, -2, 0}
	r := AddFourNumbers(A, B, C, D)
	if r == 2 {
		t.Log("正确")
	} else {
		t.Errorf("答案错误")
	}
}
