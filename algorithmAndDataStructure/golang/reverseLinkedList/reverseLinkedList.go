package reverseLinkedList

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return reverseListIteration(head)
	//return reverseListRecursion(head, nil)
}

func reverseListIteration(head *ListNode) (newHead *ListNode) {
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

func reverseListRecursion(curNode, preNode *ListNode) (newHead *ListNode) {
	if curNode == nil {
		return preNode
	}
	next := curNode.Next
	curNode.Next = preNode
	newHead = reverseListRecursion(next, curNode)
	return
}
