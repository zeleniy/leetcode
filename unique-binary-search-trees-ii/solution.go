package unique_binary_search_trees_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}

	return generateSubtrees(1, n)
}

func generateSubtrees(start, end int) []*TreeNode {

	if start > end {
		return []*TreeNode{nil}
	}

	var trees []*TreeNode

	for i := start; i <= end; i++ {

		leftTrees := generateSubtrees(start, i-1)
		rightTrees := generateSubtrees(i+1, end)

		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = leftTree
				root.Right = rightTree
				trees = append(trees, root)
			}
		}
	}

	return trees
}
