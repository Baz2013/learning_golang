package leetcode_cn

/**
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

https://leetcode-cn.com/problems/reverse-string
*/

func reverseString(s []byte) {
	low := 0
	high := len(s) - 1

	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
}
