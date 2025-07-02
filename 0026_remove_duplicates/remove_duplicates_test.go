package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int // the expected slice after removing duplicates (first k elements)
		k        int   // the expected number of unique elements
	}{
		{input: []int{1, 1, 2}, expected: []int{1, 2}, k: 2},
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4}, expected: []int{0, 1, 2, 3, 4}, k: 5},
		{input: []int{}, expected: []int{}, k: 0},
		{input: []int{1, 2, 3, 3, 3}, expected: []int{1, 2, 3}, k: 3},
	}

	for _, test := range tests {
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)

		got := removeDuplicates(inputCopy)

		if got != test.k {
			t.Errorf("removeDuplicates(%v) returned k=%d; want %d", test.input, got, test.k)
		}

		if !reflect.DeepEqual(inputCopy[:got], test.expected) {
			t.Errorf("removeDuplicates(%v) modified nums to %v; want %v", test.input, inputCopy[:got], test.expected)
		}

	}
}
