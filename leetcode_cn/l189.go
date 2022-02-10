package leetcode_cn

// 空间复杂度不是 O(1), 需额外空间
func rotate1(nums []int, k int) {
	t := []int{}
	l := len(nums) - 1
	for i := 0; i < k; i++ {
		n := nums[l-i]
		t = append([]int{n}, t...)
	}

	for i := 0; i < l-k+1; i++ {
		t = append(t, nums[i])
	}

	for i, v := range t {
		nums[i] = v
	}
}

/**
先将数组整体反转，再反转不需要轮转的数据
*/

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l // k大于数组长度的情况

	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)

}

func reverse(nums []int, i int, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
