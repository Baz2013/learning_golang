package leetcode_cn

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	data := []int{0, 1, 0, 3, 12}
	target := []int{1, 3, 12, 0, 0}
	moveZeroes(data)

	fmt.Println(data)
	if SliceValueEqual(data, target) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	data = []int{0}
	target = []int{0}
	moveZeroes(data)
	if SliceValueEqual(data, target) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	data = []int{5, 6, 0, 0, 7, 0, 0}
	target = []int{5, 6, 7, 0, 0, 0, 0}
	moveZeroes(data)
	fmt.Println(data)
	if SliceValueEqual(data, target) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	data = []int{0, 0, 2, 3, 5, 0, 0}
	target = []int{2, 3, 5, 0, 0, 0, 0}
	moveZeroes(data)
	fmt.Println(data)
	if SliceValueEqual(data, target) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

}
