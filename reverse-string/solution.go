package reverse_string

func reverseStringRecursive(s []byte) {

	n := len(s)

	if n <= 1 {
		return
	}

	tmp := s[0]
	s[0] = s[n-1]
	s[n-1] = tmp

	reverseStringRecursive(s[1 : n-1])
}

func reverseStringIterative(s []byte) {

	for i, n := 0, len(s); i < n/2; i++ {
		tmp := s[i]
		s[i] = s[n-i-1]
		s[n-i-1] = tmp
	}
}
