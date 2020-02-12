package binarySearchTreeIterator

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type treeNodeStack struct {
	treeNodes []*TreeNode
}

func (tStack *treeNodeStack) len() int {
	return len(tStack.treeNodes)
}

func (tStack *treeNodeStack) push(treeNode *TreeNode) {
	tStack.treeNodes = append(tStack.treeNodes, treeNode)
}

func (tStack *treeNodeStack) pop() (treeNode *TreeNode) {
	stackLen := len(tStack.treeNodes)
	if stackLen == 0 {
		return
	}

	n := stackLen - 1
	treeNode = tStack.treeNodes[n]
	tStack.treeNodes = tStack.treeNodes[:n]
	return
}

type BSTIterator struct {
	stack *treeNodeStack
	node  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bSTIterator := BSTIterator{}
	bSTIterator.node = root
	bSTIterator.stack = new(treeNodeStack)
	bSTIterator.stack.treeNodes = make([]*TreeNode, 0)
	return bSTIterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	var value int
	for this.HasNext() {
		for this.node != nil {
			this.stack.push(this.node)
			this.node = this.node.Left
		}
		this.node = this.stack.pop()
		value = this.node.Val
		this.node = this.node.Right
		break
	}
	return value
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.node != nil || this.stack.len() != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
