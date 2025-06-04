package climbing_stairs

var cache = make(map[int]int)

func climbStairs(n int) int {

	if n < 3 {
		return n
	}

	if cached, exists := cache[n]; exists {
		return cached
	}

	result := climbStairs(n-1) + climbStairs(n-2)
	cache[n] = result

	return result
}
