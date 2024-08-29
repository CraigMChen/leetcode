package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	nodeQueue := []*TreeNode{root}
	valQueue := []int{0}
	for len(nodeQueue) != 0 {
		var (
			nodeNext []*TreeNode
			valNext  []int
		)
		for i := 0; i < len(nodeQueue); i++ {
			node := nodeQueue[i]
			val := valQueue[i]
			if node.Left == nil && node.Right == nil {
				if node.Val+val == targetSum {
					return true
				}
				continue
			}
			if node.Left != nil {
				nodeNext = append(nodeNext, node.Left)
				valNext = append(valNext, node.Val+val)
			}
			if node.Right != nil {
				nodeNext = append(nodeNext, node.Right)
				valNext = append(valNext, node.Val+val)
			}
		}
		nodeQueue = nodeNext
		valQueue = valNext
	}
	return false
}
