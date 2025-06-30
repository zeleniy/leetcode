package remove_linked_list_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	val    int
	head   *ListNode
}

func BenchmarkRemoveElements(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			removeElements(data.head, data.val)
		}
	}
}

func TestRemoveElements(t *testing.T) {

	for _, data := range getTestDataSet() {

		head := removeElements(data.head, data.val)

		for _, val := range data.answer {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}

		assert.Nil(t, head)
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1, 2, 3, 4, 5}, 6, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}},
		{[]int{}, 1, nil},
		{[]int{}, 7, &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}}},
		{[]int{2}, 1, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{[]int{1}, 2, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
	}
}
