package leetcode_cn

import "testing"

func TestTwoSum(t *testing.T) {

	data := []int{2, 7, 11, 15}
	target := 9
	expectedResult := []int{1, 2}
	result := twoSum(data, target)
	if SliceValueEqual(expectedResult, result) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	data = []int{2, 3, 4}
	target = 6
	expectedResult = []int{1, 3}
	result = twoSum(data, target)
	if SliceValueEqual(expectedResult, result) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	data = []int{-1, 0}
	target = -1
	expectedResult = []int{1, 2}
	result = twoSum(data, target)
	if SliceValueEqual(expectedResult, result) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
