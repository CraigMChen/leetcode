package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := int(1e5)
	var preorder func(node *TreeNode, depth int)
	preorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if depth < ans {
				ans = depth
			}
			return
		}
		preorder(node.Left, depth + 1)
		preorder(node.Right, depth + 1)
	}
	preorder(root, 1)
	return ans
}
