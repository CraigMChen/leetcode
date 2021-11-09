package ac1

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func printTree(root *TreeNode) [][]string {
	depth := 0
	var getDepth func(level int, node *TreeNode)
	getDepth = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		if level > depth {
			depth = level
		}
		getDepth(level+1, node.Left)
		getDepth(level+1, node.Right)
	}
	getDepth(1, root)

	res := make([][]string, depth)
	index := 0
	var dfs func(level int, node *TreeNode)
	dfs = func(level int, node *TreeNode) {
		if node == nil {
			if level <= depth {
				index += (1<<(depth-level+1)) - 1
			}
			return
		}
		dfs(level+1, node.Left)
		if res[level-1] == nil {
			res[level-1] = make([]string, (1<<depth)-1)
		}
		res[level-1][index] = strconv.FormatInt(int64(node.Val), 10)
		index++
		dfs(level+1, node.Right)
	}
	dfs(1, root)
	return res
}
