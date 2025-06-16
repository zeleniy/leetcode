package flood_fill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer        [][]int
	sr, sc, color int
	image         [][]int
}

func BenchmarkFloodFill(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			floodFill(data.image, data.sr, data.sc, data.color)
		}
	}
}

func TestFloodFill(t *testing.T) {
	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, floodFill(data.image, data.sr, data.sc, data.color))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
		{[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, 1, 1, 2, [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}},
	}
}
