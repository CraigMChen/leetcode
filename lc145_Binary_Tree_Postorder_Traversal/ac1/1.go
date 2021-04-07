package ac1

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}
