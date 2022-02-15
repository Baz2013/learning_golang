package leetcode_cn

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {

	data := "abcabcbb"
	r := lengthOfLongestSubstring(data)
	if r == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败， r==%d", r)
	}

	data = "bbbbb"
	r = lengthOfLongestSubstring(data)
	if r == 1 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败， r==%d", r)
	}

	data = "pwwkew"
	r = lengthOfLongestSubstring(data)
	if r == 3 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败， r==%d", r)
	}

	data = " "
	r = lengthOfLongestSubstring(data)
	if r == 1 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败， r==%d", r)
	}

	data = ""
	r = lengthOfLongestSubstring(data)
	if r == 0 {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败， r==%d", r)
	}
}
