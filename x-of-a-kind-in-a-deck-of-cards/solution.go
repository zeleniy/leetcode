package x_of_a_kind_in_a_deck_of_cards

func hasGroupsSizeX(deck []int) bool {

	if len(deck) < 2 {
		return false
	}

	countMap := make(map[int]int)

	for _, num := range deck {
		countMap[num]++
	}

	gcdValue := 0

	for _, count := range countMap {
		gcdValue = gcd(gcdValue, count)
		if gcdValue == 1 {
			return false
		}
	}

	return gcdValue >= 2
}

func gcd(x, y int) int {

	for y != 0 {
		x, y = y, x%y
	}

	return x
}
