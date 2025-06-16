package binary_tree_inorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	node   *TreeNode
}

func BenchmarkInorderTraversal(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			inorderTraversal(data.node)
		}
	}
}

func TestInorderTraversal(t *testing.T) {
	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, inorderTraversal(data.node))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{}, nil},
		{[]int{1}, &TreeNode{Val: 1}},
		{[]int{1, 3, 2}, &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
	}
}
