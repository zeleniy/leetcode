package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {

	result := make([]int, len(nums))

	i := 0
	for ; i < len(nums) && nums[i] < 0; i++ {
	}

	j, k := i, 0
	i--

	for ; i >= 0 && j < len(nums); k++ {
		if nums[j]+nums[i] < 0 {
			result[k] = nums[j] * nums[j]
			j++
		} else {
			result[k] = nums[i] * nums[i]
			i--
		}
	}

	for ; i >= 0; i, k = i-1, k+1 {
		result[k] = nums[i] * nums[i]
	}

	for ; j < len(nums); j, k = j+1, k+1 {
		result[k] = nums[j] * nums[j]
	}

	return result
}
