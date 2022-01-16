package leetcode_cn

// leetcode 704

// 在列表a中查找数字n， 找到返回其下标， 找不到返回-1
func BinarySearch1(a []int, n int) int {
	left := 0
	right := len(a) - 1 // 查找区间为 a[left, right]， 一个左闭右闭的区间

	for left <= right {
		middle := (left + right) / 2
		if a[middle] == n {
			return middle
		} else if a[middle] < n {
			left = middle + 1
		} else if a[middle] > n {
			right = middle - 1
		}
	}

	return -1
}

func BinarySearch(a []int, n int) int {
	left := 0
	right := len(a) // 查找区间为 a[left, right)， 一个左闭右开的区间

	for left < right {
		middle := (left + right) / 2
		if a[middle] == n {
			return middle
		} else if a[middle] < n {
			left = middle + 1
		} else if a[middle] > n {
			right = middle
		}
	}

	return -1
}
