package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {

	charMap := make(map[byte]int)
	maxLen := 0

	for left, right, n := 0, 0, len(s); right < n; right++ {

		char := s[right]

		if index, exists := charMap[char]; exists && index >= left {
			left = index + 1
		}

		charMap[char] = right
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
