package sort_an_array

func MergeSortRecursiveWithIndexes(nums []int) []int {

	mergeSortRecursiveWithIndexes(nums, 0, len(nums)-1)
	return nums
}

func mergeSortRecursiveWithIndexes(nums []int, left, right int) {

	if left >= right {
		return
	}

	mid := (left + right) / 2

	mergeSortRecursiveWithIndexes(nums, left, mid)
	mergeSortRecursiveWithIndexes(nums, mid+1, right)

	mergeWithIndexes(nums, left, mid, right)
}

func mergeWithIndexes(nums []int, left, mid, right int) {

	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for ; i <= mid && j <= right; k++ {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
	}

	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = nums[i]
	}

	for ; j <= right; j, k = j+1, k+1 {
		tmp[k] = nums[j]
	}

	for idx := 0; idx < len(tmp); idx++ {
		nums[left+idx] = tmp[idx]
	}
}

func MergeSortRecursiveWithSlices(nums []int) []int {

	n := len(nums)

	if n <= 1 {
		return nums
	}

	mid := n / 2
	left := MergeSortRecursiveWithSlices(nums[0:mid])
	right := MergeSortRecursiveWithSlices(nums[mid:])

	return mergeWithSlices(left, right)
}

func mergeWithSlices(a, b []int) []int {

	result := make([]int, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result[i+j] = a[i]
			i++
		} else {
			result[i+j] = b[j]
			j++
		}
	}

	copy(result[i+j:], a[i:])
	copy(result[i+j:], b[j:])

	return result
}
