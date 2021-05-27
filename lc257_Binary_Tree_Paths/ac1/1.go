package ac1

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func binaryTreePaths(root *TreeNode) []string {
	var (
		getPath func(root *TreeNode, path string)
		ans []string
	)

	getPath = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			path += strconv.FormatInt(int64(root.Val), 10)
			ans = append(ans, path)
			return
		}
		path = path + strconv.FormatInt(int64(root.Val), 10) + "->"
		getPath(root.Left, path)
		getPath(root.Right, path)
	}
	getPath(root, "")
	return ans
}
