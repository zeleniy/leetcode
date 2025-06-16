package decode_string

import "strings"

func decodeString(s string) string {

	var numbersStack []int
	var stringsStack []string
	currentNumber := 0
	currentString := ""
	builder := strings.Builder{}

	for _, char := range s {
		if char >= '0' && char <= '9' {
			currentNumber = currentNumber*10 + int(char-'0')
		} else if char == '[' {
			numbersStack = append(numbersStack, currentNumber)
			stringsStack = append(stringsStack, currentString)
			currentNumber = 0
			currentString = ""
		} else if char == ']' {
			repeatCount := numbersStack[len(numbersStack)-1]
			numbersStack = numbersStack[:len(numbersStack)-1]

			previousString := stringsStack[len(stringsStack)-1]
			stringsStack = stringsStack[:len(stringsStack)-1]

			builder.Reset()
			builder.WriteString(previousString)
			for i := 0; i < repeatCount; i++ {
				builder.WriteString(currentString)
			}
			currentString = builder.String()
		} else {
			currentString += string(char)
		}
	}

	return currentString
}
