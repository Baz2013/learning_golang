package leetcode_cn

import "fmt"

/**
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}
*/

type MyLinkedList struct {
	Head *ListNode `json:"header"`
	Tail *ListNode `json:"tail"`
	Lens int       `json:"lens"`
}

func Constructor() MyLinkedList {

	return MyLinkedList{
		Head: &ListNode{},
		Tail: &ListNode{},
		Lens: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Lens {
		return -1
	}

	if index == 0 {
		return this.Head.Next.Val
	}

	head := this.Head.Next
	for head.Next != nil {
		index--
		head = head.Next
		if index == 0 {
			return head.Val
		}
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}

	cur := this.Head.Next
	this.Head.Next = node
	node.Next = cur

	if this.Lens == 0 {
		this.Tail.Next = node
	}

	this.Lens++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	if this.Lens == 0 {
		this.Tail.Next = node
		this.Head.Next = node
		this.Lens++
		return
	}
	tail := this.Tail.Next
	tail.Next = node
	this.Tail.Next = node
	this.Lens++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if this.Lens == index {
		this.AddAtTail(val)
		return
	}

	if this.Lens < index {
		return
	}

	if index < 0 {
		this.AddAtHead(val)
		return
	}

	head := this.Head
	for head != nil {
		if index == 0 {
			cur := head.Next
			head.Next = &ListNode{
				Val:  val,
				Next: cur,
			}
			this.Lens++
			break
		}
		index--
		head = head.Next
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Lens-1 {
		return
	}

	l := index

	head := this.Head
	for head != nil {

		if l == 0 {
			cur := head.Next
			head.Next = cur.Next
			if index == this.Lens-1 {
				// 删除最后一个节点时做特殊处理
				this.Tail.Next = head
			}
			this.Lens--
			break
		}
		head = head.Next
		l--
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func (this *MyLinkedList) PrintMyLinkedList() {
	head := this.Head.Next

	fmt.Println("**********************")
	fmt.Printf("(len:%d)  ", this.Lens)
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
