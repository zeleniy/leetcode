package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	tail.Next = head
	toPassBy := length - k - 1

	for i := 0; i < toPassBy; i++ {
		head = head.Next
	}

	newHead := head.Next
	head.Next = nil

	return newHead
}
