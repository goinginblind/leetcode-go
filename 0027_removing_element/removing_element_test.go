package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name       string
		inputSlice []int
		inputValue int
		expected   []int // the expected slice after removing Element (first k elements)
		k          int   // the expected number of unique elements
	}{
		{
			name:       "I Regular 1",
			inputSlice: []int{0, 1, 2, 2, 3, 0, 4, 2},
			inputValue: 2,
			expected:   []int{0, 1, 3, 0, 4},
			k:          5,
		},
		{
			name:       "II Empty slice",
			inputSlice: []int{},
			inputValue: 2,
			expected:   []int{},
			k:          0,
		},
		{
			name:       "III Slice of size 1, -rm",
			inputSlice: []int{0},
			inputValue: 0,
			expected:   []int{},
			k:          0,
		},
		{
			name:       "IV Slice of size 1, return 1",
			inputSlice: []int{0},
			inputValue: 1,
			expected:   []int{0},
			k:          1,
		},
		{
			name:       "V Regular 2 many rm'd at the end",
			inputSlice: []int{0, 1, 1, 1, 2, 3, 4, 5, 6, 1, 2, 3, 1, 1, 1, 1, 1},
			inputValue: 1,
			expected:   []int{0, 2, 3, 4, 5, 6, 2, 3},
			k:          8,
		},
		{
			name:       "VI Regular 3, single rm at the end",
			inputSlice: []int{0, 1, 2, 3, 4},
			inputValue: 4,
			expected:   []int{0, 1, 2, 3},
			k:          4,
		},
		{
			name:       "VII Regular 4, no -rm",
			inputSlice: []int{0, 1, 2, 3, 4},
			inputValue: 5,
			expected:   []int{0, 1, 2, 3, 4},
			k:          5,
		},
		{
			name:       "IIX Delete all",
			inputSlice: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			inputValue: 1,
			expected:   []int{},
			k:          0,
		},
	}

	for _, test := range tests {
		inputSliceCopy := make([]int, len(test.inputSlice))
		copy(inputSliceCopy, test.inputSlice)

		got := removeElement(inputSliceCopy, test.inputValue)

		if got != test.k {
			t.Errorf("\ncase: %v\n removeElement(%v)\n returned k=%d;\n want %v", test.name, test.inputSlice, got, test.k)
		}

		if !reflect.DeepEqual(inputSliceCopy[:got], test.expected) {
			t.Errorf("\ncase: %v\n removeElement(%v)\n modified nums to %v;\n want %v", test.name, test.inputSlice, inputSliceCopy[:got], test.expected)
		}

	}
}
