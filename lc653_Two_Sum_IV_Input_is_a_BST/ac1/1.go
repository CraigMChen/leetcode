package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 用map记录已经出现过的数
func findTarget(root *TreeNode, k int) bool {
	var (
		ans bool
		dfs func(node *TreeNode)
	)
	m := make(map[int]struct{})
	dfs = func(node *TreeNode) {
		if ans {
			return
		}
		if node == nil {
			return
		}
		dfs(node.Left)
		if _, ok := m[k - node.Val]; ok {
			ans = true
			return
		}
		m[node.Val] = struct{}{}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
