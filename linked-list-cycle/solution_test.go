package linked_list_cycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	head   *ListNode
}

func BenchmarkHasCycle(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			hasCycle(data.head)
		}
	}
}

func TestHasCycle(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, hasCycle(data.head))
	}
}

func getTestDataSet() []TestCaseDataSet {

	node11 := &ListNode{Val: 1}
	node12 := &ListNode{Val: 2}
	node13 := &ListNode{Val: 0}
	node14 := &ListNode{Val: -4}

	node11.Next = node12
	node12.Next = node13
	node13.Next = node14
	node14.Next = node12

	node21 := &ListNode{Val: 1}
	node22 := &ListNode{Val: 2}

	node21.Next = node22
	node22.Next = node21

	return []TestCaseDataSet{
		{true, node11},
		{true, node21},
		{false, &ListNode{Val: 1}},
		{false, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{false, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
	}
}
