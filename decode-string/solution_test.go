package decode_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer string
	code   string
}

func BenchmarkDecodeString(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			decodeString(data.code)
		}
	}
}

func TestDecodeString(t *testing.T) {
	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, decodeString(data.code))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{"", ""},
		{"abc", "abc"},
		{"abcbc", "a2[bc]"},
		{"aaabcbc", "3[a]2[bc]"},
		{"accaccacc", "3[a2[c]]"},
		{"abcabccdcdcdef", "2[abc]3[cd]ef"},
	}
}
