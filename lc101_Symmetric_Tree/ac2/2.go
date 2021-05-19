package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var que1, que2 []*TreeNode
	que1 = append(que1, root)
	que2 = append(que2, root)

	for len(que1) != 0 && len(que2) != 0 {
		var next1, next2 []*TreeNode
		for i := 0; i < len(que1); i++ {
			if que1[i] == nil && que2[i] == nil {
				continue
			}
			if que1[i] != nil && que2[i] != nil {
				if que1[i].Val != que2[i].Val {
					return false
				}
				next1 = append(next1, que1[i].Left, que1[i].Right)
				next2 = append(next2, que2[i].Right, que2[i].Left)
			} else {
				return false
			}
		}
		que1 = next1
		que2 = next2
	}
	return true
}
