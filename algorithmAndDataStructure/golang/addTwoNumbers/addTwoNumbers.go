package addTwoNumbers

// ListNode is definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func getValueAndNext(l *ListNode) (next *ListNode, value int) {
	value = 0
	if l != nil {
		value = l.Val
		l = l.Next
	}
	next = l
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	for carry, tmp := 0, &head; l1 != nil || l2 != nil || carry > 0; {
		var x, y int
		l1, x = getValueAndNext(l1)
		l2, y = getValueAndNext(l2)
		sum := x + y + carry
		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		*tmp = node
		tmp = &node.Next

		carry = sum / 10
	}
	return head
}
