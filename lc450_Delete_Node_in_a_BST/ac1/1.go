package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 修改节点的值
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Right != nil { // 若节点右儿子非空
			suc := successor(root)                       // 则找到它的后继节点
			root.Val = suc.Val                           // 将前驱节点的值赋给该节点
			root.Right = deleteNode(root.Right, suc.Val) // 并删除该后继节点
		} else if root.Left != nil { // 若节点左儿子非空
			pre := predecessor(root)                   // 则找到它的前驱节点
			root.Val = pre.Val                         // 将后继节点的值赋给该节点
			root.Left = deleteNode(root.Left, pre.Val) // 并删除该前驱节点
		} else { // 若节点为叶子节点
			return nil // 则直接删除
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

func successor(node *TreeNode) *TreeNode {
	suc := node.Right
	for suc.Left != nil {
		suc = suc.Left
	}
	return suc
}
