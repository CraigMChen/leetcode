package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var (
		max []int
		dfs func(node *TreeNode, cur int)
	)
	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		if cur >= len(max) {
			max = append(max, node.Val)
		}
		if node.Val > max[cur] {
			max[cur] = node.Val
		}
		dfs(node.Left, cur + 1)
		dfs(node.Right, cur + 1)
	}
	dfs(root, 0)
	return max
}
