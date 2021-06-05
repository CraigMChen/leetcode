package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	ans := &TreeNode{Val: root1.Val + root2.Val}
	que := []*TreeNode{ans}
	que1 := []*TreeNode{root1}
	que2 := []*TreeNode{root2}
	for len(que) != 0 {
		node := que[0]
		que = que[1:]
		node1 := que1[0]
		que1 = que1[1:]
		node2 := que2[0]
		que2 = que2[1:]
		left1 := node1.Left
		left2 := node2.Left
		right1 := node1.Right
		right2 := node2.Right
		if left1 != nil || left2 != nil {
			if left1 != nil && left2 != nil {
				left := &TreeNode{Val: left1.Val + left2.Val}
				node.Left = left
				que = append(que, left)
				que1 = append(que1, left1)
				que2 = append(que2, left2)
			} else if left1 != nil && left2 == nil {
				node.Left = left1
			} else {
				node.Left = left2
			}
		}
		if right1 != nil || right2 != nil {
			if right1 != nil && right2 != nil {
				right := &TreeNode{Val: right1.Val + right2.Val}
				node.Right = right
				que = append(que, right)
				que1 = append(que1, right1)
				que2 = append(que2, right2)
			} else if right1 != nil && right2 == nil {
				node.Right = right1
			} else {
			 	node.Right = right2
			}
		}
	}
	return ans
}
