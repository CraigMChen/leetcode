package ac2

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 迭代
func preorderTraversal(root *TreeNode) []int {
	var (
		ans   []int
		stack []*TreeNode
	)
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return ans
}
