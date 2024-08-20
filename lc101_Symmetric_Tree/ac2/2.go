package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	que1 := []*TreeNode{root.Left}
	que2 := []*TreeNode{root.Right}
	for len(que1) > 0 && len(que2) > 0 {
		node1 := que1[len(que1)-1]
		que1 = que1[:len(que1)-1]
		node2 := que2[len(que2)-1]
		que2 = que2[:len(que2)-1]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		que1 = append(que1, node1.Left, node1.Right)
		que2 = append(que2, node2.Right, node2.Left)
	}
	return len(que1) == len(que2)
}
