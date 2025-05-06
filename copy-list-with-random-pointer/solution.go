package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	current := head

	for current != nil {
		copy := &Node{Val: current.Val}
		copy.Next = current.Next
		current.Next = copy
		current = copy.Next
	}

	newHead := head.Next
	current = head

	for ; current != nil; current = current.Next.Next {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
	}

	current = head

	for current != nil && current.Next != nil {
		copy := current.Next
		current = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
	}

	return newHead
}
