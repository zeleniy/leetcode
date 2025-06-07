package unique_binary_search_trees

var cache = make(map[int]int)

func numTrees(n int) int {

	if n < 2 {
		return 1
	}

	if cached, exists := cache[n]; exists {
		return cached
	}

	count := 0

	for rootVal := 1; rootVal <= n; rootVal++ {
		count += numTrees(rootVal-1) * numTrees(n-rootVal)
	}

	cache[n] = count

	return count
}
