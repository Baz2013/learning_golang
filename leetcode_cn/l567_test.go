package leetcode_cn

import "testing"

func TestCheckInclusion(t *testing.T) {

	s1 := "ab"
	s2 := "eidbaooo"
	r := checkInclusion(s1, s2)
	if r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	s1 = "ab"
	s2 = "eidboaoo"
	r = checkInclusion(s1, s2)
	if !r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	s1 = "adc"
	s2 = "dcda"
	r = checkInclusion(s1, s2)
	if r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}

	s1 = "abc"
	s2 = "bbbca"
	r = checkInclusion(s1, s2)
	if r {
		t.Log("测试通过")
	} else {
		t.Errorf("测试失败")
	}
}
