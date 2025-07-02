package rotate_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	k      int
	head   *ListNode
}

func BenchmarkRotateRight(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			rotateRight(data.head, data.k)
		}
	}
}

func TestRotateRight(t *testing.T) {

	for _, data := range getTestDataSet() {

		head := rotateRight(data.head, data.k)

		for _, val := range data.answer {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}

		assert.Nil(t, head)
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{4, 5, 1, 2, 3}, 2, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}},
		{[]int{2, 0, 1}, 4, &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}},
	}
}
