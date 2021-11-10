package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func widthOfBinaryTree(root *TreeNode) int {
	var (
		levels []int
		dfs    func(level, value int, node *TreeNode)
	)
	res := 0
	dfs = func(level, value int, node *TreeNode) {
		if node == nil {
			return
		}
		if level >= len(levels) {
			levels = append(levels, value)
		}
		width := value - levels[level] + 1
		if width > res {
			res = width
		}
		dfs(level+1, value*2, node.Left)
		dfs(level+1, value*2+1, node.Right)
	}
	dfs(0, 1, root)
	return res
}
