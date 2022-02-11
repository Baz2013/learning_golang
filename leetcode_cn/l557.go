package leetcode_cn

import (
	"strings"
)

/**
Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
*/
// 该方法内存和耗时都很高
func reverseWords1(s string) string {
	space := " "
	word := []string{}
	res := ""
	items := strings.Split(s, "")
	for _, v := range items {
		if v != space {
			word = append([]string{v}, word...)
		}
		if v == " " {
			res = res + strings.Join(word, "")
			word = []string{}
			res = res + space
		}
	}

	res = res + strings.Join(word, "")

	return res
}

// 还是双指针
// 该方法极大的减少的内存使用和耗时
func reverseWords(s string) string {
	space := byte(' ')
	items := []byte(s)
	l := len(items)

	for i := 0; i < l; i++ {
		j := i
		for j < l && items[j] != space {
			j++
		}

		s := i
		e := j - 1
		for ; s < e; s, e = s+1, e-1 {
			items[s], items[e] = items[e], items[s]
		}

		i = j
	}

	return string(items)
}
