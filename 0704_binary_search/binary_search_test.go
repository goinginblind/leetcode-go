package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		tgt      int
		expected int
	}{
		{
			name:     "I Regular, odd",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			tgt:      3,
			expected: 2,
		},
		{
			name:     "II Regular, even",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			tgt:      3,
			expected: 2,
		},
		{
			name:     "III Regular, even, close to mid",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			tgt:      6,
			expected: 5,
		},
		{
			name:     "IV Target out of range",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			tgt:      14,
			expected: -1,
		},
		{
			name:     "V Target doesn't exist",
			arr:      []int{1, 2, 3, 5, 6, 7, 8, 9, 10},
			tgt:      4,
			expected: -1,
		},
		{
			name:     "VI Target doesn't exist",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11},
			tgt:      8,
			expected: -1,
		},
		{
			name:     "VII Very short",
			arr:      []int{1, 2, 3},
			tgt:      3,
			expected: 2,
		},
		{
			name:     "IIX Very short, target absent",
			arr:      []int{0, 2, 3},
			tgt:      1,
			expected: -1,
		},
		{
			name:     "IX Very short, target absent",
			arr:      []int{0},
			tgt:      1,
			expected: -1,
		},
		{
			name:     "X Very short",
			arr:      []int{0},
			tgt:      0,
			expected: 0,
		},
		{
			name:     "XI Regular",
			arr:      []int{0},
			tgt:      0,
			expected: 0,
		},
		{
			name:     "XIII Two elements, target first",
			arr:      []int{2, 4},
			tgt:      2,
			expected: 0,
		},
		{
			name:     "XIV Two elements, target second",
			arr:      []int{2, 4},
			tgt:      4,
			expected: 1,
		},
		{
			name:     "XV Two elements, target absent",
			arr:      []int{2, 4},
			tgt:      3,
			expected: -1,
		},
		{
			name:     "XVI Large array, target at start",
			arr:      []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			tgt:      10,
			expected: 0,
		},
		{
			name:     "XVII Large array, target at end",
			arr:      []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			tgt:      100,
			expected: 9,
		},
		{
			name:     "XVIII Negative numbers, target present",
			arr:      []int{-10, -5, 0, 5, 10},
			tgt:      -5,
			expected: 1,
		},
		{
			name:     "XIX Negative numbers, target absent",
			arr:      []int{-10, -5, 0, 5, 10},
			tgt:      -6,
			expected: -1,
		},
	}

	for _, test := range tests {
		got := search(test.arr, test.tgt)

		if got != test.expected {
			t.Errorf("case: %v\nexpected: %v\ngot: %v\n", test.name, test.expected, got)
		}
	}
}
