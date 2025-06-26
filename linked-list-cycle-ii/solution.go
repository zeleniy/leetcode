package linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}

	return nil
}
