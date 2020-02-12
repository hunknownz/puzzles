package binaryTreeInorderTraversal

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

func inorderTraversal(root *TreeNode) []int {
	inorderValues := make([]int, 0)
	if root == nil {
		return inorderValues
	}
	treeNodeStack := &treeNodeStack{
		treeNodes: make([]*TreeNode, 0),
	}

	node := root
	for node != nil || treeNodeStack.len() != 0 {
		for node != nil {
			treeNodeStack.push(node)
			node = node.Left
		}
		node = treeNodeStack.pop()
		inorderValues = append(inorderValues, node.Val)
		node = node.Right
	}
	return inorderValues
}
