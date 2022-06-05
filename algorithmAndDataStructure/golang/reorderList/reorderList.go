package reorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListFromMiddle(head *ListNode) (rightNode *ListNode) {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	rightNode = slow.Next
	slow.Next = nil
	return
}

func reverseLinkedList(head *ListNode) (newHead *ListNode) {
	var preNode *ListNode
	curNode := head
	for curNode != nil {
		next := curNode.Next
		curNode.Next = preNode
		newHead, preNode = curNode, curNode
		curNode = next
	}
	return
}

func mergeLeftRightHead(left, right *ListNode) {
	l, r := left, right
	for l != nil && r != nil {
		lNext := l.Next
		l.Next = r
		rNext := r.Next
		r.Next = lNext
		l = lNext
		r = rNext
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	rightNode := splitListFromMiddle(head)
	rightNode = reverseLinkedList(rightNode)
	mergeLeftRightHead(head, rightNode)
}
