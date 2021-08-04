package ac3
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// map记录父节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var (
		ans *TreeNode
		dfs func(pre, cur *TreeNode)
	)
	parents := make(map[*TreeNode]*TreeNode)
	dfs = func(pre, cur *TreeNode) {
		if cur == nil {
			return
		}
		parents[cur] = pre
		dfs(cur, cur.Left)
		dfs(cur, cur.Right)
	}
	dfs(nil, root)

	existed := make(map[*TreeNode]struct{})
	for p != nil {
		existed[p] = struct{}{}
		p = parents[p]
	}
	for q != nil {
		if _, ok := existed[q]; ok {
			ans = q
			break
		}
		q = parents[q]
	}
	return ans
}
