package two_sum

func twoSum(nums []int, target int) []int {

	numMap := make(map[int]int, 4)

	for i, num := range nums {

		if j, exists := numMap[target-num]; exists {
			return []int{j, i}
		}

		numMap[num] = i
	}

	return []int{}
}
