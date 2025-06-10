package daily_temperatures

func dailyTemperatures(temperatures []int) []int {

	length := len(temperatures)
	answer := make([]int, length)
	stack := make([]int, length)
	top := 0

	for i, temperature := range temperatures {

		for ; top > 0 && temperature > temperatures[stack[top-1]]; top-- {
			prevIndex := stack[top-1]
			answer[prevIndex] = i - prevIndex
		}

		stack[top] = i
		top++
	}

	return answer
}
