package leetcode_cn

/**
Given a string s, find the length of the longest substring without repeating characters.
*/

func lengthOfLongestSubstring(s string) int {

	window := make(map[byte]int)
	n := len(s)
	start := 0 // 滑动窗口开始
	end := 0   // 滑动窗口结束
	cnt := 0

	for end < n {
		b := s[end]
		window[b]++
		end++
		for window[b] > 1 {
			window[s[start]]--
			start++
		}

		cnt = max(end-start, cnt)

	}

	return cnt
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
