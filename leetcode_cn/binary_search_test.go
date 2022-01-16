package leetcode_cn

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 7, 9, 10}
	i := BinarySearch(a, 8)
	if i != -1 {
		t.Errorf("查找数字8失败")
	}

	i = BinarySearch(a, 2)
	if i != 1 {
		t.Errorf("查找数字2失败")
	}

	i = BinarySearch(a, 10)
	if i != 6 {
		t.Errorf("查找数字2失败")
	}

	i = BinarySearch(a, 4)
	if i != 3 {
		t.Errorf("查找数字2失败")
	}
}
