package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func sumRootToLeaf(root *TreeNode) int {
	var (
		ans int
		dfs func(node *TreeNode, cur int)
	)
	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			ans += cur << 1 + node.Val
			return
		}
		dfs(node.Left, cur << 1 + node.Val)
		dfs(node.Right, cur << 1 + node.Val)
	}
	dfs(root, 0)
	return ans
}