package main

import (
	"testing"
)

func foo(any) any {
	return nil
}

func TestFoo(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "I",
			input:    1,
			expected: 2,
		},
		{
			name:     "II",
			input:    2,
			expected: 3,
		},
	}

	for _, test := range tests {
		got := foo(test.input)

		if got != test.expected {
			t.Errorf("case: %v\nexpected: %v\ngot: %v\n", test.name, test.expected, got)
		}
	}
}
