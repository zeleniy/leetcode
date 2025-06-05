package max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {

	globalMax := 0
	localMax := 0

	for _, num := range nums {
		if num == 1 {
			localMax++
		} else {
			globalMax = max(globalMax, localMax)
			localMax = 0
		}
	}

	return max(globalMax, localMax)
}

func max(a int, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}
