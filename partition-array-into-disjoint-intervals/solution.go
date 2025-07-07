package partition_array_into_disjoint_intervals

func partitionDisjoint(nums []int) int {

	n := len(nums)
	maxLeft := nums[0]
	currMax := nums[0]
	partitionIdx := 0

	for i := 1; i < n; i++ {

		currMax = max(currMax, nums[i])

		if nums[i] < maxLeft {
			maxLeft = currMax
			partitionIdx = i
		}
	}

	return partitionIdx + 1
}
