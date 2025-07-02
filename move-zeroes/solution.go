package move_zeroes

func moveZeroes(nums []int) {

	insertTo := 0

	for i := insertTo; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[insertTo] = nums[i]
			insertTo++
		}
	}

	for i := insertTo; i < len(nums); i++ {
		nums[i] = 0
	}
}
