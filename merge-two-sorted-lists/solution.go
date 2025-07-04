package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	node := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}

		node = node.Next
	}

	if list1 != nil {
		node.Next = list1
	} else {
		node.Next = list2
	}

	return dummyHead.Next
}
