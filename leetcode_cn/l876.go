package leetcode_cn

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
使用快慢指针，慢指针一次走一步，快指针一次走两步，这样当快指针到链尾的时候，慢指针正好在中间
*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
