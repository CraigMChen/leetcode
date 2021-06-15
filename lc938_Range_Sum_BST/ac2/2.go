package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	que := []*TreeNode{root}
	for len(que) != 0 {
		node := que[0]
		que = que[1:]
		if node == nil {
			continue
		}
		if node.Val < low {
			que = append(que, node.Right)
		}  else if node.Val > high {
			que = append(que, node.Left)
		} else {
			ans += node.Val
			que = append(que, node.Left)
			que = append(que, node.Right)
		}
	}
	return ans
}
