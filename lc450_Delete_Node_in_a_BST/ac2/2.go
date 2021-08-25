package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 移动子树
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil { // 若左子树为空
			return root.Right // 则返回右子树
		} else if root.Right == nil { // 若右子树为空
			return root.Left // 则返回左子树
		} else { // 若都不为空
			pre := predecessor(root) // 则把该节点的右子树接到前驱节点的右子树上
			pre.Right = root.Right
			return root.Left // 并返回该节点的左子树
			// suc := successor(root) // 或者把该节点的左子树接到后继节点的左子树上
			// suc.Left = root.Left
			// return root.Right // 并返回该节点的右子树
		}
	}
	return root
}

func predecessor(node *TreeNode) *TreeNode {
	pre := node.Left
	for pre.Right != nil {
		pre = pre.Right
	}
	return pre
}
