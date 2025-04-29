package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {

	l := len(nums1) - 1
	m--
	n--

	for ; m >= 0 && n >= 0; l = l - 1 {
		if nums1[m] > nums2[n] {
			nums1[l] = nums1[m]
			m--
		} else {
			nums1[l] = nums2[n]
			n--
		}
	}

	for ; n >= 0; n, l = n-1, l-1 {
		nums1[l] = nums2[n]
	}
}
