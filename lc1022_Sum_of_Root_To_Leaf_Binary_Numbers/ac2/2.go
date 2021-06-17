package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	que := []*TreeNode{root}
	val := []int{0}
	for len(que) != 0 {
		node := que[0]
		que = que[1:]
		num := val[0]
		val = val[1:]
		cur := num << 1 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += cur
		}
		if node.Left != nil {
			que = append(que, node.Left)
			val = append(val, cur)
		}
		if node.Right != nil {
			que = append(que, node.Right)
			val = append(val, cur)
		}
	}
	return ans
}
