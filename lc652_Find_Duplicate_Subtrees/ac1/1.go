package ac1

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用map存储每个子树的序列化结果
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	hash := make(map[string]int)

	var (
		res []*TreeNode
		dfs func(node *TreeNode) string
	)
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "@"
		}
		s := strconv.FormatInt(int64(node.Val), 32)
		str := s + "," + dfs(node.Left) + "," + dfs(node.Right)
		hash[str]++
		if hash[str] == 2 {
			res = append(res, node)
		}
		return str
	}
	dfs(root)
	return res
}
