package search_in_a_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	answer *TreeNode
	target int
	tree   *TreeNode
}

func TestSearchBST(t *testing.T) {

	for _, dataSet := range getTestDataSet() {

		root := searchBST(dataSet.tree, dataSet.target)

		if dataSet.answer == nil {
			assert.Nil(t, root)
		} else {
			assert.Equal(t, dataSet.answer.Val, root.Val)
		}
	}
}

func getTestDataSet() []TestCase {

	return []TestCase{
		{&TreeNode{Val: 2}, 2, &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}}},
		{nil, 5, &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}}},
	}
}
