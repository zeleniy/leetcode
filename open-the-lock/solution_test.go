package open_the_lock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer   int
	target   string
	deadends []string
}

func BenchmarkOpenLock(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			openLock(data.deadends, data.target)
		}
	}
}

func TestOpenLock(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, openLock(data.deadends, data.target))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{6, "0202", []string{"0201", "0101", "0102", "1212", "2002"}},
		{1, "0009", []string{"8888"}},
		{-1, "8888", []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}},
	}
}
