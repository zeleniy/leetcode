package validate_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, min, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
}
