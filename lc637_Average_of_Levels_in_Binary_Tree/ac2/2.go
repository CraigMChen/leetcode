package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func averageOfLevels(root *TreeNode) []float64 {
	var (
		ans []float64
		counts []int
		sum []int
		dfs func(node *TreeNode, level int)
	)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(counts) {
			counts[level]++
			sum[level] += node.Val
		} else {
			counts = append(counts, 1)
			sum = append(sum, node.Val)
		}
		dfs(node.Left, level + 1)
		dfs(node.Right, level + 1)
	}
	dfs(root, 0)
	for i := 0; i < len(counts); i++ {
		ans = append(ans, float64(sum[i]) / float64(counts[i]))
	}
	return ans
}
