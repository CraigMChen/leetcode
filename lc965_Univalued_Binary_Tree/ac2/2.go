package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isLeftUnival := root.Left == nil || root.Left.Val == root.Val && isUnivalTree(root.Left)
	isRightUnival := root.Right == nil || root.Right.Val == root.Val && isUnivalTree(root.Right)
	return isLeftUnival && isRightUnival
}
