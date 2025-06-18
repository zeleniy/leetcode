package validate_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	root   *TreeNode
}

func BenchmarkIsValidBST(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			isValidBST(data.root)
		}
	}
}

func TestIsValidBST(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, isValidBST(data.root))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{true, nil},
		{true, &TreeNode{Val: -2147483648}},
		{true, &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
		{false, &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}},
		{false, &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}},
		{false, &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}},
	}
}
