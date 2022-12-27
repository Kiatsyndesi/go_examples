package Valid_Parentheses

import "strings"

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(value string) {
	*s = append(*s, value)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		lastIndex := len(*s) - 1
		elemToPop := (*s)[lastIndex]
		*s = (*s)[:lastIndex]

		return elemToPop, true
	}
}

func ValidParentheses(parens string) bool {

	var stack stack
	sliceOfParens := strings.Split(parens, "")

	for _, paren := range sliceOfParens {
		if paren == "(" {
			stack.push(paren)
		} else {
			if stack.isEmpty() {
				return false
			} else {
				topParen := stack[len(stack)-1]

				if paren == ")" && topParen == "(" {
					stack.pop()
				} else {
					return false
				}
			}
		}
	}

	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}
