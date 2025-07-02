package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "I",
			input:    "()",
			expected: true,
		},
		{
			name:     "II",
			input:    "(){}[]",
			expected: true,
		},
		{
			name:     "III",
			input:    "([{}])",
			expected: true,
		},
		{
			name:     "IV",
			input:    "({})[]",
			expected: true,
		},
		{
			name:     "V",
			input:    "(){}]",
			expected: false,
		},
		{
			name:     "VI",
			input:    "({)}[]",
			expected: false,
		},
	}

	for _, test := range tests {
		got := isValid(test.input)

		if got != test.expected {
			t.Errorf("case: %v\nexpected: %v\ngot: %v\n", test.name, test.expected, got)
		}
	}
}
