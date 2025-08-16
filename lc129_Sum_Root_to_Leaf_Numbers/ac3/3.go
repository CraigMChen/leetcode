package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, cur int) int {
	if node == nil {
		return 0
	}
	sum := cur*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return dfs(node.Left, sum) + dfs(node.Right, sum)
}
