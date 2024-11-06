package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	expressions := []string{
		"{[()]}",
		"{[(])}",
		"{{[[(())]]}}",
	}

	for _, exp := range expressions {
		fmt.Printf("Expression: %s, is valid: %v\n", exp, isValid(exp))
	}
}
