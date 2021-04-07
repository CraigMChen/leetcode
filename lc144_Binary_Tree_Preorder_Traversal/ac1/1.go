package ac1

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}
