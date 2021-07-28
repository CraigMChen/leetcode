package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{root.Val}
	level := []*TreeNode{root}
	for len(level) != 0 {
		var next []*TreeNode
		for _, node := range level {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next= append(next, node.Right)
			}
		}
		if next != nil {
			ans = append(ans, next[len(next) - 1].Val)
		}
		level = next
	}
	return ans
}