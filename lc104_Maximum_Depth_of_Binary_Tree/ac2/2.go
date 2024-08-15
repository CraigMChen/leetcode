package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
// 按层数搜索
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	que := []*TreeNode{root}
	depth := 0
	for len(que) > 0 {
		for num := len(que); num > 0; num-- {
			node := que[0]
			que = que[1:]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		depth++
	}
	return depth
}
