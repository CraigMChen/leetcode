package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:   val,
			Left:  root,
		}
	}
	nodeQue := []*TreeNode{root}
	depthQue := []int{1}
	for len(nodeQue) != 0 {
		node := nodeQue[0]
		nodeQue = nodeQue[1:]
		cur := depthQue[0]
		depthQue = depthQue[1:]
		if cur == depth-1 {
			node.Left = &TreeNode{
				Val:  val,
				Left: node.Left,
			}
			node.Right = &TreeNode{
				Val:   val,
				Right: node.Right,
			}
			continue
		}
		if node.Left != nil {
			nodeQue = append(nodeQue, node.Left)
			depthQue = append(depthQue, cur+1)
		}
		if node.Right != nil {
			nodeQue = append(nodeQue, node.Right)
			depthQue = append(depthQue, cur+1)
		}
	}
	return root
}
