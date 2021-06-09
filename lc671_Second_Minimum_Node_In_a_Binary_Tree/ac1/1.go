package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func findSecondMinimumValue(root *TreeNode) int {
	min := root.Val
	var ans *int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val > min && (ans == nil || node.Val < *ans) {
			ans = &node.Val
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	if ans == nil {
		return -1
	}
	return *ans
}
