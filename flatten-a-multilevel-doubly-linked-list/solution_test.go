package flatten_a_multilevel_doubly_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	answer []int
	head   *Node
}

func TestFlatten(t *testing.T) {

	for _, testCase := range getTestDataSet() {

		head := flatten(testCase.head)

		for _, val := range testCase.answer {
			assert.Equal(t, val, head.Val)
			head = head.Next
		}
	}
}

func getTestDataSet() []TestCase {

	return []TestCase{
		{[]int{1}, &Node{Val: 1}},
		{[]int{1, 2}, &Node{Val: 1, Next: &Node{Val: 2}}},
		{[]int{1, 2}, &Node{Val: 1, Child: &Node{Val: 2}}},
		// {[]int{1, 2, 3, 7, 8, 9, 10, 4, 5, 6}, &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Child: &Node{Val: 7, Next: &Node{Val: 8, Next: &Node{Val: 9, Next: &Node{Val: 10}}}}, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6}}}}}}},
		{[]int{1, 2, 3}, &Node{Val: 1, Next: &Node{Val: 2, Child: &Node{Val: 3}}}},
		{[]int{1, 3, 2}, &Node{Val: 1, Next: &Node{Val: 2}, Child: &Node{Val: 3}}},
	}
}
