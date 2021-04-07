package ac1

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var next []*TreeNode
		var cur []int
		for i := 0; i < len(queue); i++ {
			cur = append(cur, queue[i].Val)
			if queue[i].Left != nil {
				next = append(next, queue[i].Left)
			}
			if queue[i].Right != nil {
				next = append(next, queue[i].Right)
			}
		}
		res = append(res, cur)
		queue = next
	}
	return res
}