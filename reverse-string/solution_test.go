package reverse_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []byte
	word   []byte
}

func TestReverseStringRecursive(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		reverseStringRecursive(dataSet.word)
		testReverseString(t, dataSet)
	}
}

func TestReverseStringIterative(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		reverseStringIterative(dataSet.word)
		testReverseString(t, dataSet)
	}
}

func testReverseString(t *testing.T, dataSet TestCaseDataSet) {

	assert.Equal(t, len(dataSet.answer), len(dataSet.word))

	for i := range dataSet.word {
		assert.Equal(t, dataSet.answer[i], dataSet.word[i])
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]byte{'o', 'l', 'l', 'e', 'h'}, []byte{'h', 'e', 'l', 'l', 'o'}},
		{[]byte{'h', 'a', 'n', 'n', 'a', 'H'}, []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
	}
}
