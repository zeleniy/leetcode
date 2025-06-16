package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	result := []int{}

	if root == nil {
		return result
	}

	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		current = current.Right
	}

	return result
}
