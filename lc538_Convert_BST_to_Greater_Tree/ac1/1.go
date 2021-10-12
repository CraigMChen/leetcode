package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先求和 再中序遍历
func convertBST(root *TreeNode) *TreeNode {
	var (
		getSum func(node *TreeNode)
		inorder func(node *TreeNode)
		sum int
	)
	getSum = func(node *TreeNode) {
		if node == nil {
			return
		}
		sum += node.Val
		getSum(node.Left)
		getSum(node.Right)
	}
	getSum(root)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		val := node.Val
		node.Val = sum
		sum -= val
		inorder(node.Right)
	}
	inorder(root)
	return root
}