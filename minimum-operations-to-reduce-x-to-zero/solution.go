package minimum_operations_to_reduce_x_to_zero

func minOperations(nums []int, x int) int {

	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	target := totalSum - x

	if target < 0 {
		return -1
	}

	n := len(nums)

	if target == 0 {
		return n
	}

	maxLen := -1
	currentSum := 0

	for left, right := 0, 0; right < n; right++ {

		currentSum += nums[right]

		for currentSum > target && left <= right {
			currentSum -= nums[left]
			left++
		}

		if currentSum == target {
			maxLen = max(maxLen, right-left+1)
		}
	}

	if maxLen == -1 {
		return -1
	}

	return n - maxLen
}
