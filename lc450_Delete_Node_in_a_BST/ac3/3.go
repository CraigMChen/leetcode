package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
// 移动子树的迭代版
func deleteNode(root *TreeNode, key int) *TreeNode {
	var targetNode, parent *TreeNode
	for targetNode = root; targetNode != nil && targetNode.Val != key; {
		parent = targetNode
		if targetNode.Val > key {
			targetNode = targetNode.Left
		} else {
			targetNode = targetNode.Right
		}
	}
	if targetNode == nil {
		return root
	}
	if targetNode.Left == nil && targetNode.Right == nil {
		targetNode = nil
	} else if targetNode.Left == nil {
		targetNode = targetNode.Right
	} else if targetNode.Right == nil {
		targetNode = targetNode.Left
	} else {
		suc := successor(targetNode)
		suc.Left = targetNode.Left
		targetNode = targetNode.Right
	}
	if parent == nil {
		return targetNode
	}
	if parent.Left != nil && parent.Left.Val == key {
		parent.Left = targetNode
	} else {
		parent.Right = targetNode
	}
	return root
}

func successor(node *TreeNode) *TreeNode {
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node
}
