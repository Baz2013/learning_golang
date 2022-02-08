package lcof

import "sort"

// FindRepeatNumber 排序的方式
func FindRepeatNumber(n []int) int {

	sort.Ints(n) // go 语言中的切片排序 https://tonybai.com/2020/11/26/slice-sort-in-go/
	for i, v := range n {
		if i == 0 {
			continue
		}
		if n[i-1] == v {
			return v
		}
	}

	return -1
}

// FindRepeatNumber1 哈希表的方式
func FindRepeatNumber1(n []int) int {

	var m = map[int]int{}
	for _, v := range n {
		k := m[v]
		if k == 1 {
			return v
		} else {
			m[v] = 1
		}
	}

	return -1
}

// FindRepeatNumber2 原地交换的方式
func FindRepeatNumber2(n []int) int {
	for i, v := range n {
		if i == v {
			continue
		}

		if n[v] == v {
			return v
		}

		t := n[v]
		n[v] = v
		n[i] = t
	}

	return -1
}
