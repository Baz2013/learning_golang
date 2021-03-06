package leetcode_cn

import (
	"fmt"
	"testing"
)

func TestDesignLinkedList(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtHead(2)
	obj.AddAtHead(3)
	obj.AddAtHead(4)
	obj.PrintMyLinkedList()
	param_1 := obj.Get(1)
	fmt.Println(param_1)
	obj.AddAtTail(5)
	obj.AddAtTail(6)
	obj.AddAtTail(7)
	obj.PrintMyLinkedList()
	obj.AddAtIndex(3, 8)
	obj.PrintMyLinkedList()
	obj.AddAtIndex(0, 9)
	obj.PrintMyLinkedList()
	obj.AddAtIndex(-1, 10)
	obj.PrintMyLinkedList()
	obj.DeleteAtIndex(0)
	obj.PrintMyLinkedList()
	obj.DeleteAtIndex(1)
	obj.PrintMyLinkedList()
}

func TestDesignLinkedList2(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	linkedList.PrintMyLinkedList()
	val := linkedList.Get(1) //返回2
	fmt.Println(val)
	linkedList.PrintMyLinkedList()
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	linkedList.PrintMyLinkedList()
	val = linkedList.Get(1) //返回3
	fmt.Println(val)

}

func TestDesignLinkedList3(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	linkedList.PrintMyLinkedList()
	val := linkedList.Get(1) //返回2
	fmt.Println(val)
	linkedList.PrintMyLinkedList()
	linkedList.DeleteAtIndex(0) //现在链表是2-> 3
	linkedList.PrintMyLinkedList()
	val = linkedList.Get(0) //返回2
	fmt.Println(val)

}

func TestDesignLinkedList4(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtTail(1)
	val := linkedList.Get(0) //返回2
	fmt.Println(val)
	linkedList.AddAtTail(2)
	linkedList.PrintMyLinkedList()
	linkedList.AddAtTail(3) //链表变为1-> 2-> 3
	linkedList.PrintMyLinkedList()
	linkedList.AddAtIndex(1, 4)
	linkedList.PrintMyLinkedList()
	//val := linkedList.Get(1) //返回2
	//fmt.Println(val)
	//linkedList.PrintMyLinkedList()
	//linkedList.DeleteAtIndex(0) //现在链表是2-> 3
	//linkedList.PrintMyLinkedList()
	//val = linkedList.Get(0) //返回2
	//fmt.Println(val)

}

func TestDesignLinkedList5(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(84)
	obj.AddAtTail(2)
	obj.AddAtTail(39)
	obj.Get(3)
	obj.Get(1)
	obj.AddAtTail(42)
	obj.AddAtIndex(1, 80)
	obj.AddAtHead(14)
	obj.AddAtHead(1)
	obj.AddAtTail(53)
	obj.AddAtTail(98)
	obj.AddAtTail(19)
	obj.AddAtTail(12)
	obj.PrintMyLinkedList()
	obj.Get(2)
	obj.AddAtHead(16)
	obj.AddAtHead(33)
	obj.PrintMyLinkedList()
	obj.AddAtIndex(4, 17)
	obj.PrintMyLinkedList()
	obj.AddAtIndex(6, 8)
	obj.PrintMyLinkedList()
	obj.AddAtHead(37)
	obj.PrintMyLinkedList()
	obj.AddAtTail(43)
	obj.PrintMyLinkedList()
	obj.DeleteAtIndex(11)
	obj.PrintMyLinkedList() //ok, len 16
	obj.AddAtHead(80)
	obj.AddAtHead(31)
	obj.AddAtIndex(13, 23)
	obj.PrintMyLinkedList()
	obj.AddAtTail(17)
	obj.PrintMyLinkedList() // ok , len 20
	obj.Get(4)
	obj.AddAtIndex(10, 0)
	obj.AddAtTail(21)
	obj.AddAtHead(73)
	obj.AddAtHead(22)
	obj.AddAtIndex(24, 37)
	obj.AddAtTail(14)
	obj.AddAtHead(97)
	obj.AddAtHead(8)
	obj.PrintMyLinkedList() // ok, 28
	val := obj.Get(6)
	fmt.Println(val)
	obj.DeleteAtIndex(17)
	obj.PrintMyLinkedList()
	obj.AddAtTail(50)
	obj.AddAtTail(28)
	obj.AddAtHead(76)
	obj.AddAtTail(79)
	obj.PrintMyLinkedList()
	obj.Get(18)
	obj.DeleteAtIndex(30)
	obj.PrintMyLinkedList()
	obj.AddAtTail(5)
	obj.AddAtHead(9)
	obj.AddAtTail(83)
	obj.DeleteAtIndex(3)
	obj.PrintMyLinkedList()
	obj.AddAtTail(40)
	obj.DeleteAtIndex(26)
	obj.PrintMyLinkedList()
	obj.AddAtIndex(20, 90)
	obj.PrintMyLinkedList()
	obj.DeleteAtIndex(30) //////
	obj.AddAtTail(40)
	obj.AddAtHead(56)
	obj.AddAtIndex(15, 23)
	obj.AddAtHead(51)
	obj.AddAtHead(21)
	obj.Get(26)
	obj.AddAtHead(83)
	obj.Get(30)
	obj.AddAtHead(12)
	obj.DeleteAtIndex(8)
	obj.Get(4)
	obj.AddAtHead(20)
	obj.AddAtTail(45)
	obj.Get(10)
	obj.AddAtHead(56)
	obj.Get(18)
	obj.AddAtTail(33)
	obj.Get(2)
	obj.AddAtTail(70)
	obj.AddAtHead(57)
	obj.AddAtIndex(31, 24)
	obj.AddAtIndex(16, 92)
	obj.AddAtHead(40)
	obj.AddAtHead(23)
	obj.DeleteAtIndex(26)
	obj.Get(1)
	obj.AddAtHead(92)
	obj.AddAtIndex(3, 78)
	obj.AddAtTail(42)
	obj.Get(18)
	obj.AddAtIndex(39, 9)
	obj.Get(13)
	obj.AddAtIndex(33, 17)
	obj.Get(51)
	obj.AddAtIndex(18, 95)
	obj.AddAtIndex(18, 33)
	obj.AddAtHead(80)
	obj.AddAtHead(21)
	obj.AddAtTail(7)
	obj.AddAtIndex(17, 46)
	obj.Get(33)
	obj.AddAtHead(60)
	obj.AddAtTail(26)
	obj.AddAtTail(4)
	obj.AddAtHead(9)
	obj.Get(45)
	obj.AddAtTail(38)
	obj.AddAtHead(95)
	obj.AddAtTail(78)
	obj.Get(54)
	obj.AddAtIndex(42, 86)
}
