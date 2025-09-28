package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归版中序遍历
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func isValidBST(root *TreeNode) bool {
	prev := -1<<31 - 1
	var inOrder func(*TreeNode) bool
	inOrder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		l := inOrder(node.Left)
		if node.Val <= prev {
			return false
		}
		prev = node.Val
		return l && inOrder(node.Right)
	}
	return inOrder(root)
}
