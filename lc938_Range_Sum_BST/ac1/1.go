package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Val < low {
			return dfs(node.Right)
		}
		if node.Val > high {
			return dfs(node.Left)
		}
		return node.Val + dfs(node.Left) + dfs(node.Right)
	}
	return dfs(root)
}
