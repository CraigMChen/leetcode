package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func rightSideView(root *TreeNode) []int {
	var (
		ans []int
		dfs func(node *TreeNode, level int)
	)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level >= len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, level + 1)
		dfs(node.Left, level + 1)
	}
	dfs(root, 0)
	return ans
}