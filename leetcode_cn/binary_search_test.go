package leetcode_cn

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 7, 9, 10}
	i := BinarySearch(a, 8)
	//fmt.Println(i)
	if i != -1 {
		t.Errorf("查找数字8失败")
	}

	i = BinarySearch(a, 2)
	//fmt.Println(i)
	if i != 1 {
		t.Errorf("查找数字2失败")
	}

	i = BinarySearch(a, 10)
	//fmt.Println(i)
	if i != 6 {
		t.Errorf("查找数字10失败")
	}

	i = BinarySearch(a, 4)
	//fmt.Println(i)
	if i != 3 {
		t.Errorf("查找数字4失败")
	}

	i = BinarySearch(a, 1)
	//fmt.Println(i)
	if i != 0 {
		t.Errorf("查找数字1失败")
	}

	i = BinarySearch(a, 9)
	//fmt.Println(i)
	if i != 5 {
		t.Errorf("查找数字9失败")
	}

	i = BinarySearch(a, 7)
	//fmt.Println(i)
	if i != 4 {
		t.Errorf("查找数字7失败")
	}

	i = BinarySearch(a, 11)
	//fmt.Println(i)
	if i != -1 {
		t.Errorf("查找数字11失败")
	}
}

func TestBinarySearch1(t *testing.T) {
	a := []int{1, 2, 3, 4, 7, 9, 10}
	i := BinarySearch1(a, 8)
	//fmt.Println(i)
	if i != -1 {
		t.Errorf("查找数字8失败")
	}

	i = BinarySearch1(a, 2)
	//fmt.Println(i)
	if i != 1 {
		t.Errorf("查找数字2失败")
	}

	i = BinarySearch1(a, 10)
	//fmt.Println(i)
	if i != 6 {
		t.Errorf("查找数字10失败")
	}

	i = BinarySearch1(a, 4)
	//fmt.Println(i)
	if i != 3 {
		t.Errorf("查找数字4失败")
	}

	i = BinarySearch1(a, 1)
	//fmt.Println(i)
	if i != 0 {
		t.Errorf("查找数字1失败")
	}

	i = BinarySearch1(a, 9)
	//fmt.Println(i)
	if i != 5 {
		t.Errorf("查找数字9失败")
	}

	i = BinarySearch1(a, 7)
	//fmt.Println(i)
	if i != 4 {
		t.Errorf("查找数字7失败")
	}

	i = BinarySearch1(a, 11)
	//fmt.Println(i)
	if i != -1 {
		t.Errorf("查找数字11失败")
	}
}

func TestBinarySearch2(t *testing.T) {
	a := []int{1, 2, 3, 4, 7, 9, 10}
	i := BinarySearch2(a, 8)
	//fmt.Println(i)
	if i != -1 {
		t.Errorf("查找数字8失败")
	}

	i = BinarySearch2(a, 2)
	//fmt.Println(i)
	if i != 1 {
		t.Errorf("查找数字2失败")
	}

	i = BinarySearch2(a, 10)
	//fmt.Println(i)
	if i != 6 {
		t.Errorf("查找数字10失败")
	}

	i = BinarySearch2(a, 4)
	//fmt.Println(i)
	if i != 3 {
		t.Errorf("查找数字4失败")
	}

	i = BinarySearch2(a, 1)
	//fmt.Println(i)
	if i != 0 {
		t.Errorf("查找数字1失败")
	}

	i = BinarySearch2(a, 9)
	//fmt.Println(i)
	if i != 5 {
		t.Errorf("查找数字9失败")
	}

	i = BinarySearch2(a, 7)
	//fmt.Println(i)
	if i != 4 {
		t.Errorf("查找数字7失败")
	}

	i = BinarySearch2(a, 11)
	//fmt.Println(i)
	if i != -1 {
		t.Errorf("查找数字11失败")
	}
}
