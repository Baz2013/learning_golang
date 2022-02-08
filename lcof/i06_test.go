package lcof

import (
	"fmt"
	"testing"
)

func TestReversePrintValue(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	res := ReversePrintValue(l)

	fmt.Println(res)
	if len(res) == 3 && res[0] == 2 && res[1] == 3 && res[2] == 1 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
