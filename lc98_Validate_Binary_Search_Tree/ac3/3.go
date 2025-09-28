package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代版中序遍历
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func isValidBST(root *TreeNode) bool {
	prev := -1<<31 - 1
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= prev {
			return false
		}
		prev = root.Val
		root = root.Right
	}
	return true
}
