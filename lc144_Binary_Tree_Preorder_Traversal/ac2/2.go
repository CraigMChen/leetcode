package ac2

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 迭代
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		for node != nil {
			ans = append(ans, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			node = node.Left
		}
	}
	return ans
}
