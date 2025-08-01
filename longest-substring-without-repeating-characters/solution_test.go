package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	s      string
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			lengthOfLongestSubstring(data.s)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, lengthOfLongestSubstring(data.s))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{3, "abcabcbb"},
		{1, "bbbbb"},
		{3, "pwwkew"},
		{1, " "},
		{2, "au"},
		{2, "abba"},
	}
}
