package leetcode_cn

import "testing"

func TestReverseString(t *testing.T) {

	data := []byte{'h', 'e', 'l', 'l', 'o'}
	expectedResult := []byte{'o', 'l', 'l', 'e', 'h'}

	reverseString(data)
	if ByteSliceValueEqual(data, expectedResult) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	data = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	expectedResult = []byte{'h', 'a', 'n', 'n', 'a', 'H'}

	reverseString(data)
	if ByteSliceValueEqual(data, expectedResult) {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
