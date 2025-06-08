package reverse_string

func reverseString(s []byte) {

	for i, n := 0, len(s); i < n/2; i++ {
		tmp := s[i]
		s[i] = s[n-i-1]
		s[n-i-1] = tmp
	}
}
