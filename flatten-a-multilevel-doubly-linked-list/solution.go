package flatten_a_multilevel_doubly_linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {

	node := root
	for node != nil {

		if node.Child == nil {
			node = node.Next
			continue
		}

		next := node.Next
		node.Next = flatten(node.Child)
		node.Child = nil
		node.Next.Prev = node
		for node.Next != nil {
			node = node.Next
		}

		if next != nil {
			next.Prev = node
		}

		node.Next = next
	}

	return root
}
