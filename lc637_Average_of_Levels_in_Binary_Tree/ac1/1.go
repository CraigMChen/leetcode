package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
	que := []*TreeNode{root}
	for len(que) != 0 {
		var next []*TreeNode
		count := len(que)
		sum := 0.0
		for i := 0; i < count; i++ {
			cur := que[i]
			sum += float64(cur.Val)
			if cur.Left != nil {
				next = append(next, cur.Left)
			}
			if cur.Right != nil {
				next = append(next, cur.Right)
			}
		}
		ans = append(ans, sum / float64(count))
		que = next
	}
	return ans
}
