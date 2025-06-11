package evaluate_reverse_polish_notation

import (
	"strconv"
)

func evalRPN(tokens []string) int {

	stack := make([]int, (len(tokens)/2)+1)
	top := -1

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":

			b := stack[top]
			top--
			a := stack[top]

			var res int

			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}

			stack[top] = res

		default:
			num, _ := strconv.Atoi(token)
			top++
			stack[top] = num
		}
	}

	return stack[0]
}
