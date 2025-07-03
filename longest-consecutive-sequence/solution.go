package longest_consecutive_sequence

func longestConsecutive(nums []int) int {

	n := len(nums)

	if n < 2 {
		return n
	}

	numsSet := make(map[int]bool)

	for _, num := range nums {
		numsSet[num] = true
	}

	longestStreak := 0

	for num := range numsSet {
		if !numsSet[num-1] {

			streakLength := 1
			nextNum := num + 1

			for numsSet[nextNum] {
				streakLength++
				nextNum++
			}

			if streakLength > longestStreak {
				longestStreak = streakLength
			}
		}
	}

	return longestStreak
}
