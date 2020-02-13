package reverseLinkedListII

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	startHead, oldHead := findStartHead(head, m)
	newHead, oldTail := reverseLinkedListIteration(startHead, m, n)
	if oldHead != nil {
		oldHead.Next = newHead
		newHead = head
	}
	startHead.Next = oldTail
	return newHead
}

func reverseLinkedListIteration(curNode *ListNode, m, n int) (newHead, oldTail *ListNode) {
	var preNode *ListNode
	cnt := m
	for curNode != nil && cnt <= n {
		next := curNode.Next
		curNode.Next = preNode
		preNode, newHead = curNode, curNode
		curNode = next
		cnt++
	}
	oldTail = curNode
	return
}

func findStartHead(head *ListNode, m int) (newStartHead, oldHead *ListNode) {
	cnt := 1
	for ; head != nil && cnt < m; cnt++ {
		if cnt == m-1 {
			oldHead = head
		}
		head = head.Next
	}
	newStartHead = head
	return
}
