package odd_even_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	head   *ListNode
}

func BenchmarkRemoveElements(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			oddEvenList(data.head)
		}
	}
}

func TestRemoveElements(t *testing.T) {

	for _, data := range getTestDataSet() {

		head := oddEvenList(data.head)

		for _, val := range data.answer {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}

		assert.Nil(t, head)
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1, 3, 2}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
		{[]int{1, 3, 5, 2, 4}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}},
		{[]int{2, 3, 6, 7, 1, 5, 4}, &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}}}}},
	}
}
