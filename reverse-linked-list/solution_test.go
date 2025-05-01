package reverse_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	expected []int
	head     *ListNode
}

func TestSearch(t *testing.T) {

	for _, dataSet := range getSearchData() {
		head := reverseList(dataSet.head)
		for _, val := range dataSet.expected {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}
	}
}

func getSearchData() []TestCase {

	return []TestCase{
		{[]int{4, 3, 2, 1}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}},
	}
}
