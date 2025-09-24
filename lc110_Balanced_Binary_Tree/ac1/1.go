package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自顶向上的递归
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 &&
		isBalanced(root.Right) && isBalanced(root.Left)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
