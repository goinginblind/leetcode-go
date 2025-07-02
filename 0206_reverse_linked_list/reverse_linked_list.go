package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// LL example: [1 -> 2 -> 3] ->> [3 -> 2 -> 1]

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	for current := head; current != nil; {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
