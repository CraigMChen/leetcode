package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		positiveOrder [][]int
		ans [][]int
	)
	que := []*TreeNode{root}
	for len(que) != 0 {
		var (
			next []*TreeNode
			levelOrder []int
		)
		for i := range que {
			levelOrder = append(levelOrder, que[i].Val)
			if que[i].Left != nil {
				next = append(next, que[i].Left)
			}
			if que[i].Right != nil {
				next = append(next, que[i].Right)
			}
		}
		que = next
		positiveOrder = append(positiveOrder, levelOrder)
	}
	for i := len(positiveOrder) - 1; i >= 0; i-- {
		ans = append(ans, positiveOrder[i])
	}
	return ans
}
