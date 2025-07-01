package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	head := dummyHead
	carry := 0

	for l1 != nil && l2 != nil {

		val := l1.Val + l2.Val + carry
		carry = val / 10
		val %= 10

		l1.Val = val
		head.Next = l1
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	l := l1
	if l == nil {
		l = l2
	}

	for l != nil {

		val := l.Val + carry
		carry = val / 10
		val %= 10

		l.Val = val
		head.Next = l
		head = head.Next
		l = l.Next
	}

	if carry == 1 {
		head.Next = &ListNode{Val: 1}
	}

	return dummyHead.Next
}
