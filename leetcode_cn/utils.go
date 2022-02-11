package leetcode_cn

// SliceValueEqual 两个切片中的值是否相同
func SliceValueEqual(a, b []int) bool {
	la := len(a)
	lb := len(b)

	if la != lb {
		return false
	}

	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func ByteSliceValueEqual(a, b []byte) bool {
	la := len(a)
	lb := len(b)

	if la != lb {
		return false
	}

	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
