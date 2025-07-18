package ac2

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 迭代
func postorderTraversal(root *TreeNode) []int {
	var (
		ans   []int
		stack []*TreeNode
		prev  *TreeNode
	)
	for root != nil || len(stack) != 0 {
		for root != nil { // 后序遍历的顺序为 左 右 中
			stack = append(stack, root) // root为中节点，不能遍历
			root = root.Left            // 先入栈
		}
		root = stack[len(stack)-1] // 出栈
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev { // 如果root没有右儿子，或root的右儿子已经遍历过了
			ans = append(ans, root.Val) // 当root没有右儿子时，root作为左节点遍历；当root右儿子已经遍历过时，root作为中节点遍历
			prev = root                 // 记录当前被遍历的节点
			root = nil                  // 防止root再次入栈
		} else { // 如果root有右儿子，则root为中节点，不能遍历
			stack = append(stack, root) // 将root入栈
			root = root.Right           // 找root的右儿子，作为下一次循环的中节点
		}
	}
	return ans
}
