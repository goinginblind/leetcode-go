package main

import "fmt"

/*Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.*/

type MinStack struct {
	stack []entry // stack has (value, minvalue) pairs
}

type entry struct {
	val int
	min int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]entry, 0),
	}
}

func (ms *MinStack) Push(val int) {
	min := val
	if len(ms.stack) > 0 && ms.stack[len(ms.stack)-1].min < val {
		min = ms.stack[len(ms.stack)-1].min
	}

	ms.stack = append(ms.stack, entry{val, min})
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1].val
}

func (ms *MinStack) GetMin() int {
	return ms.stack[len(ms.stack)-1].min
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(-1)
	obj.Push(0)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Push(3)
	obj.Push(-4)
	obj.Pop()
	obj.Push(7)
	fmt.Println(obj.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
