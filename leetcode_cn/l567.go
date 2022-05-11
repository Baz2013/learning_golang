package leetcode_cn

func checkInclusion(s1 string, s2 string) bool {

	n := len(s1)
	start := 0
	end := start + n

	for end < len(s2)+1 {
		s := s2[start:end]
		if s == s1 {
			return true
		}
		if Reverse(s) == s1 {
			return true
		}
		start++
		end++
	}
	return false
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
