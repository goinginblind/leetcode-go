package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	dummyHead *ListNode
	tail      *ListNode
	size      int
}

func Constructor() MyLinkedList {
	dummy := &ListNode{Val: -1, Next: nil}
	return MyLinkedList{
		dummyHead: dummy,
		tail:      dummy,
		size:      0,
	}
}

func (ll *MyLinkedList) Get(index int) int {
	curr := ll.dummyHead
	for range index {
		if curr == nil {
			return -1
		}
		curr = curr.Next
	}
	if curr != nil && curr.Next != nil {
		return curr.Next.Val
	} else {
		return -1
	}
}

func (ll *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val, Next: ll.dummyHead.Next}
	ll.dummyHead.Next = newNode
	if ll.size == 0 {
		ll.tail = newNode
	}
	ll.size++
}

func (ll *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val, Next: nil}
	ll.tail.Next = newNode
	ll.tail = newNode
	ll.size++
}

func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > ll.size {
		return
	}
	curr := ll.dummyHead
	for range index {
		curr = curr.Next
	}
	newNode := &ListNode{Val: val, Next: curr.Next}
	curr.Next = newNode
	if index == ll.size {
		ll.tail = newNode
	}
	ll.size++
}

func (ll *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= ll.size {
		return
	}
	curr := ll.dummyHead
	for range index {
		curr = curr.Next
	}
	if curr.Next == ll.tail {
		curr.Next = nil
		ll.tail = curr
	} else {
		curr.Next = curr.Next.Next
	}
	ll.size--
}

func main() {
	ll := Constructor()
	ll.AddAtHead(1)
	ll.AddAtTail(3)
	ll.AddAtIndex(1, 2)    // linked list becomes 1->2->3
	fmt.Println(ll.Get(1)) // return 2
	ll.DeleteAtIndex(1)    // now the linked list is 1->3
	fmt.Println(ll.Get(1)) // return 3
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
