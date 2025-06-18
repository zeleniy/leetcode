package keys_and_rooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	rooms  [][]int
}

func BenchmarkCanVisitAllRooms(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			canVisitAllRooms(data.rooms)
		}
	}
}

func TestCanVisitAllRooms(t *testing.T) {
	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, canVisitAllRooms(data.rooms))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{true, [][]int{{1}, {1}}},
		{true, [][]int{{1}, {2}, {3}, {}}},
		{false, [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
	}
}
