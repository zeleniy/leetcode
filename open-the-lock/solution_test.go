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

func TestOpenLock(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, openLock(dataSet.deadends, dataSet.target))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{6, "0202", []string{"0201", "0101", "0102", "1212", "2002"}},
		{1, "0009", []string{"8888"}},
		{-1, "8888", []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}},
	}
}
