package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge two linked lists recursively.
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursive(list2.Next, list1)
		return list2
	}
}

// Merge two linked lists iteratively.
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
