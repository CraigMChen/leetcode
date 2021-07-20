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
	for root != nil || len(stack) != 0 {
		for root != nil { // 前序遍历的顺序为 中 左 右
			ans = append(ans, root.Val) // root为中节点，先遍历
			stack = append(stack, root) // 再入栈
			root = root.Left // 不断找左儿子
		}
		root = stack[len(stack)-1] // 出栈
		stack = stack[:len(stack)-1]
		root = root.Right // 找右儿子，作为下一次循环的中节点
	}
	return ans
}
