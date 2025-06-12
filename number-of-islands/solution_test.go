package number_of_islands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	grid   [][]byte
}

func BenchmarkNumIslandsDfs(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			numIslandsDfs(data.grid)
		}
	}
}

func TestNumIslandsDfs(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, numIslandsDfs(data.grid))
	}
}

func BenchmarkNumIslandsBfs(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			numIslandsBfs(data.grid)
		}
	}
}

func TestNumIslandsBfs(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, numIslandsBfs(data.grid))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{1, [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}},
		{3, [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}},
		{3, [][]byte{
			{'1', '0', '1', '1', '0', '1', '1'},
		}},
		{1, [][]byte{
			{'1', '1', '1', '1', '1', '1', '1'},
		}},
	}
}
