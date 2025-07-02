package main

import "fmt"

// It came out not circular at all so need to revisit in about 5 days
type MyCircularQueue struct {
	queue []int
}

func Constructor(k int) MyCircularQueue {
	q := make([]int, 0, k)
	return MyCircularQueue{queue: q}
}

func (cq *MyCircularQueue) EnQueue(value int) bool {
	if len(cq.queue) == cap(cq.queue) {
		return false
	}
	cq.queue = append(cq.queue, value)
	return true
}

func (cq *MyCircularQueue) DeQueue() bool {
	if len(cq.queue) == 0 {
		return false
	}
	copy(cq.queue, cq.queue[1:])
	cq.queue = cq.queue[:len(cq.queue)-1]
	return true
}

func (cq *MyCircularQueue) Front() int {
	if len(cq.queue) == 0 {
		return -1
	}
	return cq.queue[0]
}

func (cq *MyCircularQueue) Rear() int {
	if len(cq.queue) == 0 {
		return -1
	}
	return cq.queue[len(cq.queue)-1]
}

func (cq *MyCircularQueue) IsEmpty() bool {
	return len(cq.queue) == 0
}

func (cq *MyCircularQueue) IsFull() bool {
	return len(cq.queue) == cap(cq.queue)
}

func main() {
	mcq := Constructor(3)
	mcq.EnQueue(1) // return True
	mcq.EnQueue(2) // return True
	mcq.EnQueue(3) // return True
	mcq.EnQueue(4) // return False
	mcq.Rear()     // return 3
	fmt.Println(mcq.queue)
	mcq.IsFull() // return True
	fmt.Println(mcq.queue)
	mcq.DeQueue() // return True
	mcq.DeQueue() // return True
	mcq.DeQueue() // return True
	fmt.Println(mcq.queue)
	mcq.EnQueue(4)              // return True
	mcq.EnQueue(5)              // return True
	mcq.EnQueue(6)              // return True
	fmt.Println(mcq.EnQueue(6)) // return False
	fmt.Println(mcq.queue)
	mcq.Rear() // return 4
}
