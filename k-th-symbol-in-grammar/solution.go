package k_th_symbol_in_grammar

func kthGrammar(n int, k int) int {

	if k == 1 {
		return 0
	}

	midLength := 1 << (n - 2)

	if k <= midLength {
		return kthGrammar(n-1, k)
	} else {
		return 1 - kthGrammar(n-1, k-midLength)
	}
}
