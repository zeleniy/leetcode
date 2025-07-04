package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	fast := head

	for fast != nil && fast.Next != nil {

		head = head.Next
		fast = fast.Next.Next

		if head == fast {
			return true
		}
	}

	return false
}
