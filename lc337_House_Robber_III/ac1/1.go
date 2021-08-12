package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 动态规划+递归
// 设f(x)表示选择x节点后能获得的最大值，g(x)表示不选择x节点能获得的最大值
// l，r分别为x的左儿子和右儿子，则
// f(x) = g(l) + g(r)
// g(x) = max(f(l), g(l)) + max(f(r), g(r))
func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (f, g int)
	dfs = func(node *TreeNode) (f, g int) {
		if node == nil {
			return 0, 0
		}
		fl, gl := dfs(node.Left)
		fr, gr := dfs(node.Right)
		f = node.Val + gl + gr
		g = max(fl, gl) + max(fr, gr)
		return
	}
	f, g := dfs(root)
	return max(f, g)
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
