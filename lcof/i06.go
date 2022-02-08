package lcof

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
*/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReversePrintValue(l *ListNode) []int {
	res := []int{}

	for l != nil {
		res = append([]int{l.Val}, res...)
		l = l.Next
	}

	return res
}
