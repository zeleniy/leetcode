package palindrome_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	head   *ListNode
}

func BenchmarkIsPalindromeWithFullScan(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			isPalindromeWithFullScan(data.head)
		}
	}
}

func TestIsPalindromeWithFullScan(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, isPalindromeWithFullScan(data.head))
	}
}

func BenchmarkIsPalindromeWithTwoPointers(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			isPalindromeWithTwoPointers(data.head)
		}
	}
}

func TestIsPalindromeWithTwoPointers(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, isPalindromeWithTwoPointers(data.head))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{true, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
		{false, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{false, &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}}}}},
		{false, &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
		{true, &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
	}
}
