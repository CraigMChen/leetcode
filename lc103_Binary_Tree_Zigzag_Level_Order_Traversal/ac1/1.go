package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	que := []*TreeNode{root}
	reverse := false
	for len(que) != 0 {
		var (
			next []*TreeNode
			curOrder []int
		)
		for i := len(que) - 1; i >= 0; i-- {
			node := que[i]
			curOrder = append(curOrder, node.Val)
			if reverse {
				if node.Right != nil {
					next = append(next, node.Right)
				}
				if node.Left != nil {
					next = append(next, node.Left)
				}
			} else {
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
			}
		}
		res = append(res, curOrder)
		que = next
		reverse = !reverse
	}
	return res
}
