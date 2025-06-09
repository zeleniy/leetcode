package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	s      string
}

func TestIsValid(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, isValid(dataSet.s))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{true, ""},
		{false, ")"},
		{true, "()"},
		{true, "()[]{}"},
		{false, "(]"},
		{true, "([])"},
	}
}
