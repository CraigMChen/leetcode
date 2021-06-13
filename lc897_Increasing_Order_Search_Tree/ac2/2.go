package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 修改原树
func increasingBST(root *TreeNode) *TreeNode {
	var (
		ans, pre *TreeNode
		dfs func(node *TreeNode)
	)
	ans = new(TreeNode)
	pre = ans
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		node.Left = nil
		pre.Right = node
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	return ans.Right
}
