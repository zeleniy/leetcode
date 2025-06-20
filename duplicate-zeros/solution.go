package duplicate_zeros

func duplicateZeros(arr []int) {

	n := len(arr)
	zeroCount := 0

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			zeroCount++
		}
	}

	left, right := n-1, n-1+zeroCount
	for left >= 0 {
		if right < n {
			arr[right] = arr[left]
		}
		if arr[left] == 0 {
			right--
			if right < n {
				arr[right] = 0
			}
		}
		left--
		right--
	}
}
