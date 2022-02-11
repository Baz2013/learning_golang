package leetcode_cn

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/

func moveZeroes(nums []int) {
	start := 0
	for n := 0; n < len(nums); n++ {
		if nums[n] == 0 {
			start = n
			break
		}
	}

	i := start // 从前往后扫描, 指向第一个0的位置
	j := start // 从前往后扫描, 指向第一个非零的位置

	for i < len(nums) && j < len(nums) {
		if nums[j] == 0 {
			j++
			continue
		}

		if nums[i] != 0 {
			i++
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}
