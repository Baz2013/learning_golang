package leetcode_cn

// 278. 第一个错误的版本

// mock
func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}

	return false
}

func firstBadVersion(n int) int {

	left, right := 1, n

	for left < right {
		mid := (left + right) >> 1
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
