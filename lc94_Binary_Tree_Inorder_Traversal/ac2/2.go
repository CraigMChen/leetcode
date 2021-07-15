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
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ans = append(ans, node.Val)
		node = node.Right
	}
	return ans
}
