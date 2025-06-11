package evaluate_reverse_polish_notation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	tokens []string
}

func BenchmarkEvalRPN(b *testing.B) {

	dataSet := getTestDataSet()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			evalRPN(data.tokens)
		}
	}
}

func TestEvalRPN(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, evalRPN(dataSet.tokens))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{1, []string{"1", "0", "+"}},
		{3, []string{"2", "1", "0", "+", "+"}},
		{9, []string{"2", "1", "+", "3", "*"}},
		{6, []string{"4", "13", "5", "/", "+"}},
		{22, []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
	}
}
