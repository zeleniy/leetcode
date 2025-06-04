package fibonacci_number

var cache = make(map[int]int)

func fib(n int) int {

	if n < 2 {
		return n
	}

	if cached, exists := cache[n]; exists {
		return cached
	}

	result := fib(n-1) + fib(n-2)
	cache[n] = result

	return result
}
