package ac2

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 迭代
func postorderTraversal(root *TreeNode) []int {
	var (
		ans []int
		stack []*TreeNode
		prev *TreeNode
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return ans
}
