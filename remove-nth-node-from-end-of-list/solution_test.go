package remove_nth_node_from_end_of_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	answer []int
	n      int
	list   *ListNode
}

func BenchmarkRemoveNthFromEnd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			removeNthFromEnd(data.list, data.n)
		}
	}
}

func TestRemoveNthFromEnd(t *testing.T) {

	for _, data := range getTestDataSet() {

		head := removeNthFromEnd(data.list, data.n)

		for _, val := range data.answer {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}
	}
}

func getTestDataSet() []TestCase {

	return []TestCase{
		{[]int{}, 1, &ListNode{Val: 1}},
		{[]int{1}, 1, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{[]int{2}, 2, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{[]int{1, 2, 3, 5}, 2, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}},
	}
}
