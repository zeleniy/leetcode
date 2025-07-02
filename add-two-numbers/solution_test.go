package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	l1     *ListNode
	l2     *ListNode
}

func BenchmarkAddTwoNumbers(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			addTwoNumbers(data.l1, data.l2)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {

	for _, data := range getTestDataSet() {

		head := addTwoNumbers(data.l1, data.l2)

		for _, val := range data.answer {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}

		assert.Nil(t, head)
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{7, 0, 8}, &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}},
		{[]int{6, 3}, &ListNode{Val: 4, Next: &ListNode{Val: 3}}, &ListNode{Val: 2}},
		{[]int{4}, &ListNode{Val: 2}, &ListNode{Val: 2}},
		{[]int{0}, &ListNode{Val: 0}, &ListNode{Val: 0}},
		{[]int{8, 9, 9, 9, 0, 0, 0, 1}, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}},
	}
}
