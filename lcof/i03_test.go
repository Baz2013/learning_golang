package lcof

import (
	"fmt"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	testData := []int{2, 3, 1, 0, 2, 5, 3}
	r := FindRepeatNumber(testData)

	fmt.Println(r)

	if r == 2 || r == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	testData2 := []int{1, 4, 3, 0, 2, 5, 3}
	r2 := FindRepeatNumber(testData2)

	fmt.Println(r2)

	if r2 == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}

func TestFindRepeatNumber1(t *testing.T) {
	testData := []int{2, 3, 1, 0, 2, 5, 3}
	r := FindRepeatNumber1(testData)

	fmt.Println(r)

	if r == 2 || r == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	testData2 := []int{1, 4, 3, 0, 2, 5, 3}
	r2 := FindRepeatNumber1(testData2)

	fmt.Println(r2)

	if r2 == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}

func TestFindRepeatNumber2(t *testing.T) {
	testData := []int{2, 3, 1, 0, 2, 5, 3}
	r := FindRepeatNumber2(testData)

	fmt.Println(r)

	if r == 2 || r == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	testData2 := []int{1, 4, 3, 0, 2, 5, 3}
	r2 := FindRepeatNumber2(testData2)

	fmt.Println(r2)

	if r2 == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
