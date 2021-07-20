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
		for root != nil { // 中序遍历的顺序为 左 中 右
			stack = append(stack, root) // root为中节点，不能先遍历，所以入栈
			root = root.Left // 不断找左儿子
		}
		root = stack[len(stack) - 1] // 出栈
		stack = stack[:len(stack) - 1]
		ans = append(ans, root.Val) // 当前root可以作为左节点或中节点遍历
		root = root.Right // 找root的右儿子，作为下一次循环的中节点
	}
	return ans
}
