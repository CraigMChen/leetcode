package ac2

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var ans []string
	nodeQueue := []*TreeNode{root}
	pathQueue := []string{""}
	for i := 0; i < len(nodeQueue); i++ {
		node := nodeQueue[i]
		path := pathQueue[i]
		if node.Left == nil && node.Right == nil {
			path += strconv.FormatInt(int64(node.Val), 10)
			ans = append(ans, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path + strconv.FormatInt(int64(node.Val), 10) + "->")
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path + strconv.FormatInt(int64(node.Val), 10) + "->")
		}
	}
	return ans
}
