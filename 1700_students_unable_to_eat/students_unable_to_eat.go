package main

import "fmt"

type Queue[T any] struct {
	data     []T
	head     int
	tail     int
	count    int
	capacity int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data:     make([]T, size),
		capacity: size,
	}
}

func (q *Queue[T]) Enqueue(val T) bool {
	if q.count == q.capacity {
		return false
	}
	q.data[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity
	q.count++
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.count == 0 {
		return zero, false
	}

	val := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.count--
	return val, true
}

func countStudents(students []int, sandwiches []int) int {
	queue := NewQueue[int](len(students))
	for _, student := range students {
		queue.Enqueue(student)
	}
	didNotEatCount, currSandwich := 0, 0
	for queue.count > 0 {
		if stu, _ := queue.Dequeue(); stu != sandwiches[currSandwich] {
			didNotEatCount++
			queue.Enqueue(stu)
		} else {
			didNotEatCount = 0
			currSandwich++
		}
		if didNotEatCount == queue.count {
			break
		}
	}
	return queue.count
}

func main() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	fmt.Println(countStudents(students, sandwiches))
}
