package leetcode_cn

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	m := removeNthFromEnd(head, 2)
	travleLinkedList(m)
	//if m.Val == 4 {
	//	t.Log("通过测试")
	//} else {
	//	t.Errorf("测试失败, m.Val: %d ", m.Val)
	//}

	head = &ListNode{
		Val:  1,
		Next: nil,
	}

	m = removeNthFromEnd(head, 1)
	travleLinkedList(m)

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	m = removeNthFromEnd(head, 1)
	travleLinkedList(m)

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	m = removeNthFromEnd(head, 2)
	travleLinkedList(m)

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	m = removeNthFromEnd(head, 3)
	travleLinkedList(m)
}
