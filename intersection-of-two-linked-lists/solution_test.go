package intersection_of_two_linked_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer *ListNode
	headA  *ListNode
	headB  *ListNode
}

func BenchmarkGetIntersectionNode(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			getIntersectionNode(data.headA, data.headB)
		}
	}
}

func TestGetIntersectionNode(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, getIntersectionNode(data.headA, data.headB))
	}
}

func getTestDataSet() []TestCaseDataSet {

	tail1 := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	head11 := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: tail1}}
	head12 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: tail1}}}

	tail2 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	head21 := &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: tail2}}}
	head22 := &ListNode{Val: 4, Next: tail2}

	head31 := &ListNode{Val: 1, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	head32 := &ListNode{Val: 1, Next: &ListNode{Val: 5}}

	return []TestCaseDataSet{
		{tail1, head11, head12},
		{tail2, head21, head22},
		{nil, head31, head32},
	}
}
