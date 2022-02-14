package leetcode_cn

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
双指针， 快指针比慢指针先走N步
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := dummy, dummy

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func travleLinkedList(head *ListNode) {

	fmt.Println("**********************")
	for head != nil {
		fmt.Printf("%d", head.Val)
		head = head.Next
		if head != nil {
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()
	fmt.Println("**********************")

}
