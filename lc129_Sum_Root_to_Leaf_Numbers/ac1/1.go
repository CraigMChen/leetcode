package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func sumNumbers(root *TreeNode) int {
	var (
		ans int
		dfs func(node *TreeNode, cur int)
	)
	dfs = func(node *TreeNode,cur int) {
		if node == nil {
			return
		}
		val := cur * 10 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += val
			return
		}
		dfs(node.Left, val)
		dfs(node.Right, val)
	}
	dfs(root, 0)
	return ans
}
