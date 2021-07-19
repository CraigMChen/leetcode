package ac2

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 迭代
func inorderTraversal(root *TreeNode) []int {
	var (
		ans []int
		stack []*TreeNode
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}
