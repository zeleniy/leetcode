package top_k_frequent_elements

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {

	if k == 0 {
		return []int{}
	}

	length := len(nums)

	if length == 0 {
		return []int{}
	}

	counts := make(map[int]int)

	for i := 0; i < length; i++ {
		counts[nums[i]]++
	}

	pairs := make([][2]int, 0, len(counts))

	for number, freq := range counts {
		pairs = append(pairs, [2]int{number, freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		result = append(result, pairs[i][0])
	}

	return result
}
