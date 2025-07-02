package main

func isValid(s string) bool {
	bracketPairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, char := range s {
		if _, isOpening := bracketPairs[char]; isOpening {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if char != bracketPairs[top] {
				return false
			}
		}
	}

	return len(stack) == 0
}
