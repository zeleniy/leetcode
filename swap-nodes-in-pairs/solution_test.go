package swap_nodes_in_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	expected []int
	head     *ListNode
}

func TestSwapPairs(t *testing.T) {

	for _, dataSet := range getSearchData() {
		head := swapPairs(dataSet.head)
		for _, val := range dataSet.expected {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}
	}
}

func getSearchData() []TestCase {

	return []TestCase{
		{[]int{2, 1, 4, 3}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}},
		{[]int{}, &ListNode{}},
		{[]int{1}, &ListNode{Val: 1}},
		{[]int{2, 1, 3}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
	}
}
