package valid_parentheses

var braceMap = map[byte]byte{
	'{': '}',
	'[': ']',
	'(': ')',
}

func isValid(s string) bool {

	stack := []byte{}

	for i := 0; i < len(s); i++ {

		brace := s[i]

		if closingBrace, exists := braceMap[brace]; exists {
			stack = append(stack, closingBrace)
			continue
		}

		top := len(stack) - 1

		if top < 0 || brace != stack[top] {
			return false
		}

		stack = stack[:top]
	}

	return len(stack) == 0
}
