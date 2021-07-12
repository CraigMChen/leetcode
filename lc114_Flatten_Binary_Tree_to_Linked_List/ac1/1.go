package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归+寻找前驱
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	if root.Left != nil {
		left := root.Left
		node := root.Left
		for node.Right != nil {
			node = node.Right
		}
		right := root.Right
		root.Right = left
		node.Right = right
		root.Left = nil
	}
	flatten(root.Right)
}
