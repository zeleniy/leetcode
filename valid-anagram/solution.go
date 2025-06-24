package valid_anagram

func isAnagram(s string, t string) bool {

	n := len(s)
	m := len(t)

	if n != m {
		return false
	}

	counter := [26]int{}

	for i := 0; i < n; i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}

	for _, count := range counter {
		if count != 0 {
			return false
		}
	}

	return true
}
