package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 深搜过程中遇到每层的第一个即为该层最左边的节点
func findBottomLeftValue(root *TreeNode) int {
	var (
		res int
		dfs func(node *TreeNode, cur int)
	)
	depth := -1
	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		if cur > depth {
			depth = cur
			res = node.Val
		}
		dfs(node.Left, cur + 1)
		dfs(node.Right, cur + 1)
	}
	dfs(root, 0)
	return res
}
