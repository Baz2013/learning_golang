package lcof

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {

	data := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	r := FindNumberIn2DArray(data, 20)
	if !r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	r = FindNumberIn2DArray(data, 10)
	if r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	r = FindNumberIn2DArray(data, 5)
	if r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	r = FindNumberIn2DArray(data, 31)
	if !r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
