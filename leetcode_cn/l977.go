package leetcode_cn

import "sort"

func sortedSquares(nums []int) []int {

	res := []int{}

	for _, n := range nums {
		res = append(res, n*n)
	}

	sort.Ints(res)

	return res
}
