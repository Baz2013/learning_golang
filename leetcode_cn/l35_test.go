package leetcode_cn

import "testing"

func TestSearchInsert(t *testing.T) {

	data := []int{1, 3, 5, 6}
	target := 5
	r := searchInsert(data, target)
	if r == 2 {
		t.Log("通过测试")
	} else {
		t.Errorf("测试失败， 结果： r == %d ", r)
	}

	data = []int{1, 3, 5, 6}
	target = 2
	r = searchInsert(data, target)
	if r == 1 {
		t.Log("通过测试")
	} else {
		t.Errorf("测试失败， 结果： r == %d ", r)
	}

	data = []int{1, 3, 5, 6}
	target = 7
	r = searchInsert(data, target)
	if r == 4 {
		t.Log("通过测试")
	} else {
		t.Errorf("测试失败， 结果： r == %d ", r)
	}

	data = []int{1, 3, 5, 6}
	target = 0
	r = searchInsert(data, target)
	if r == 0 {
		t.Log("通过测试")
	} else {
		t.Errorf("测试失败， 结果： r == %d ", r)
	}

	data = []int{1}
	target = 0
	r = searchInsert(data, target)
	if r == 0 {
		t.Log("通过测试")
	} else {
		t.Errorf("测试失败， 结果： r == %d ", r)
	}

	data = []int{1, 3, 5, 6}
	target = 6
	r = searchInsert(data, target)
	if r == 3 {
		t.Log("通过测试")
	} else {
		t.Errorf("测试失败， 结果： r == %d ", r)
	}

}
