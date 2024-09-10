package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 通过 dfs 求出以每个节点为根节点的子树中，以该节点为起点的最大路径和
// 一边深搜一边维护最大路径和
func maxPathSum(root *TreeNode) int {
	ans := -1000
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		sum := node.Val + max(0, left) + max(0, right) // 维护以当前节点为根的子树的最大路径和
		ans = max(ans, sum)
		return node.Val + max(0, max(left, right)) // 返回以当前节点为起点的最大路径和
	}
	dfs(root)
	return ans
}
