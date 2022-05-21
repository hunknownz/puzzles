package swapNodesInPairs

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	return swapPairsRecursion(head)
}

func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairsRecursion(next.Next)
	next.Next = head
	return next
}

func swapPairsIteration(head *ListNode) (newHead *ListNode) {
	newHead = head
	if head != nil && head.Next != nil {
		newHead = head.Next
	} else {
		return
	}

	curNode, next := head, head.Next
	nNext := next.Next
	next.Next, curNode.Next = curNode, nNext
	preNode := curNode
	curNode = nNext
	for curNode != nil && curNode.Next != nil {
		next = curNode.Next
		nNext := next.Next
		next.Next, curNode.Next = curNode, nNext
		preNode.Next = next
		preNode = curNode
		curNode = nNext
	}

	return
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println()
		return
	}

	curNode := head
	fmt.Printf("%d", curNode.Val)
	curNode = curNode.Next
	for curNode != nil {
		fmt.Printf(" -> %v", curNode.Val)
		curNode = curNode.Next
	}

	fmt.Println()
}

func main() {
	testList := &ListNode{Val: 1, Next: nil}
	testList.Next = &ListNode{Val: 2, Next: nil}
	next := testList.Next
	next.Next = &ListNode{Val: 3, Next: nil}
	next = next.Next
	next.Next = &ListNode{Val: 4, Next: nil}
	next = next.Next
	next.Next = &ListNode{Val: 5, Next: nil}
	next = next.Next
	next.Next = &ListNode{Val: 6, Next: nil}

	printList(testList)
	fmt.Printf("\n")
	testList = swapPairs(testList)
	printList(testList)
}
