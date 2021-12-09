package leetcode_cn

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	//want := ListNode{
	//	Val: 0,
	//	Next: nil,
	//}

	input := ListNode{
		Val:  0,
		Next: nil,
	}
	//if got := addTwoNumbers(&input, &input); got != &want {
	//
	//}
	got := addTwoNumbers(&input, &input)
	if got.Val != 0 {
		t.Errorf("error")
	}
}
