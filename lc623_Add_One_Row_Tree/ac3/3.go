package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜 迭代
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	var (
		stack []*TreeNode
		level []int
	)
	node := root
	cur := 1
	for len(stack) != 0 || node != nil {
		for node != nil {
			left := node.Left
			right := node.Right
			if cur == depth-1 {
				node.Left = &TreeNode{
					Val:  val,
					Left: left,
				}
				node.Right = &TreeNode{
					Val:   val,
					Right: right,
				}
				break
			}
			stack = append(stack, node)
			level = append(level, cur)
			node = left
			cur++
		}
		if len(stack) == 0 {
			break
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
		cur = level[len(level)-1]
		level = level[:len(level)-1]
		cur++
	}
	return root
}
