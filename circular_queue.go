package main

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

func (q *Queue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue[T]) IsFull() bool {
	return q.count == q.capacity
}

func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.count == 0 {
		return zero, false
	}

	return q.data[q.head], true
}
