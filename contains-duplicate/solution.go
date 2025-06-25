package contains_duplicate

func containsDuplicate(nums []int) bool {

	counter := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, exists := counter[nums[i]]; exists {
			return true
		} else {
			counter[nums[i]] = true
		}
	}

	return false
}
