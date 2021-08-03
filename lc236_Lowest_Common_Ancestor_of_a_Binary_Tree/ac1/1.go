package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var (
		ans *TreeNode
		dfs func(node *TreeNode) bool
	)
	// 表示求出当前节点为根的子树中是否包含p或q
	dfs = func(node *TreeNode) bool {
		if node == nil || ans != nil { // 剪枝
			return false
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l && r || (node == p || node == q) && (l || r) {
			ans = node
		}
		return l || r || node == p || node == q
	}
	dfs(root)
	return ans
}
