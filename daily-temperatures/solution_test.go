package daily_temperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer       []int
	temperatures []int
}

func BenchmarkDailyTemperatures(b *testing.B) {

	dataSet := getTestDataSet()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			dailyTemperatures(data.temperatures)
		}
	}
}

func TestDailyTemperatures(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, dailyTemperatures(dataSet.temperatures))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1, 1, 4, 2, 1, 1, 0, 0}, []int{73, 74, 75, 71, 69, 72, 76, 73}},
		{[]int{1, 1, 1, 0}, []int{30, 40, 50, 60}},
		{[]int{1, 1, 0}, []int{30, 60, 90}},
	}
}
