package leetcode_cn

/**
Given a 1-indexed(下标从1开始的) array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted

*/

func twoSum(numbers []int, target int) []int {
	res := []int{}

	low := 0
	high := len(numbers) - 1

	for low < high {
		if numbers[low]+numbers[high] > target {
			high--
			continue
		}

		if numbers[low]+numbers[high] < target {
			low++
			continue
		}

		if numbers[low]+numbers[high] == target {
			res = append(res, low+1)
			res = append(res, high+1)
			break
		}
	}

	return res
}
