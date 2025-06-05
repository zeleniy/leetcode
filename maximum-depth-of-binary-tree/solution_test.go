package maximum_depth_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	root   *TreeNode
}

func TestMaxDepth(t *testing.T) {

	for _, dataSet := range getMergeData() {
		assert.Equal(t, dataSet.answer, maxDepth(dataSet.root))
	}
}

func getMergeData() []TestCaseDataSet {

	return []TestCaseDataSet{
		{0, nil},
		{1, &TreeNode{}},
		{2, &TreeNode{Left: nil, Right: &TreeNode{}}},
		{2, &TreeNode{Left: &TreeNode{}, Right: nil}},
		{3, &TreeNode{Left: &TreeNode{}, Right: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}}},
	}
}
