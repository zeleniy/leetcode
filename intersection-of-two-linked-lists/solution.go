package intersection_of_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	lengthA := len(headA)
	lengthB := len(headB)

	if lengthA > lengthB {
		headA = forward(headA, lengthA-lengthB)
	} else {
		headB = forward(headB, lengthB-lengthA)
	}

	for headA != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func forward(head *ListNode, n int) *ListNode {

	for ; n > 0; n-- {
		head = head.Next
	}

	return head
}

func len(head *ListNode) int {

	length := 0

	for head != nil {
		head = head.Next
		length++
	}

	return length
}
