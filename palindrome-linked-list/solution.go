package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindromeWithTwoPointers(head *ListNode) bool {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return isEqual(reverse(slow), head)
}

func isPalindromeWithFullScan(head *ListNode) bool {

	count := 0

	for head1 := head; head1 != nil; head1 = head1.Next {
		count++
	}

	half := count / 2
	mid := head

	for ; half > 0; half-- {
		mid = mid.Next
	}

	return isEqual(reverse(mid), head)
}

func isEqual(head1 *ListNode, head2 *ListNode) bool {

	for head1 != nil {

		if head1.Val != head2.Val {
			return false
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
