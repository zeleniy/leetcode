package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	curr := head
	for curr != nil {
		clone := &Node{Val: curr.Val, Next: curr.Next}
		curr.Next = clone
		curr = clone.Next
	}

	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	curr = head
	newHead := head.Next

	for curr != nil {
		clone := curr.Next
		curr.Next = clone.Next
		if clone.Next != nil {
			clone.Next = clone.Next.Next
		}
		curr = curr.Next
	}

	return newHead
}
