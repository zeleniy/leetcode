package valid_anagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	s      string
	t      string
}

func BenchmarkIsAnagram(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			isAnagram(data.s, data.t)
		}
	}
}

func TestIsAnagram(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, isAnagram(data.s, data.t))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{true, "anagram", "nagaram"},
		{false, "rat", "car"},
	}
}
