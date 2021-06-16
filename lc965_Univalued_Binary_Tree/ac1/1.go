package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 遍历整棵树，查看节点值是否都是相同的
func isUnivalTree(root *TreeNode) bool {
	var (
		value *int
		dfs func(node *TreeNode)
	)
	ans := true
	dfs = func(node *TreeNode) {
		if !ans {
			return
		}
		if node == nil {
			return
		}
		if value == nil {
			value = &node.Val
		}
		if node.Val != *value {
			ans = false
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
