package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	answer []int
	a      *ListNode
	b      *ListNode
}

func TestAddTwoNumbers(t *testing.T) {

	for _, dataSet := range getTestDataSet() {

		head := mergeTwoLists(dataSet.a, dataSet.b)

		for _, val := range dataSet.answer {
			assert.Equal(t, val, head.Val)
			assert.True(t, true)
			head = head.Next
		}
	}
}

func getTestDataSet() []TestCase {

	return []TestCase{
		{[]int{1, 1, 2, 3, 4, 4}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		{[]int{}, nil, nil},
		{[]int{0}, nil, &ListNode{Val: 0}},
		{[]int{1, 1, 3, 4}, &ListNode{Val: 1}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
	}
}
