package x_of_a_kind_in_a_deck_of_cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	deck   []int
}

func BenchmarkHasGroupsSizeX(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			hasGroupsSizeX(data.deck)
		}
	}
}

func TestHasGroupsSizeX(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, hasGroupsSizeX(data.deck))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{true, []int{1, 2, 3, 4, 4, 3, 2, 1}},
		{false, []int{1, 1, 1, 2, 2, 2, 3, 3}},
		{false, []int{1}},
		{true, []int{1, 1, 2, 2, 2, 2}},
	}
}
