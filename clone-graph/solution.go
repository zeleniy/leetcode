package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return node
	}

	clones := make(map[int]*Node)
	clones[node.Val] = &Node{Val: node.Val}

	stack := make([]*Node, 1, 2)
	stack[0] = node

	for len(stack) > 0 {

		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		clone := clones[current.Val]

		if current.Neighbors != nil {
			clone.Neighbors = make([]*Node, len(current.Neighbors))
			for i, neighbor := range current.Neighbors {
				if _, exists := clones[neighbor.Val]; !exists {
					clones[neighbor.Val] = &Node{Val: neighbor.Val}
					stack = append(stack, neighbor)
				}
				clone.Neighbors[i] = clones[neighbor.Val]
			}
		}
	}

	return clones[node.Val]
}
