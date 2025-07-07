package subarray_product_less_than_k

func numSubarrayProductLessThanK(nums []int, k int) int {

	if k <= 1 {
		return 0
	}

	product := 1
	left := 0
	count := 0

	for right, num := range nums {

		product *= num

		for product >= k {
			product /= nums[left]
			left++
		}

		count += right - left + 1
	}

	return count
}
