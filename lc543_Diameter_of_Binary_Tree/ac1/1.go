package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 求每个节点的左右儿子为根的树的深度的和的最大值
// 树的深度求法：计算从该节点往下走，最多能走几个节点（包括它自己）
func diameterOfBinaryTree(root *TreeNode) int {
	var (
		ans int
		getDepth func(node *TreeNode) int
	)
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := getDepth(node.Left)
		rightDepth := getDepth(node.Right)
		ans = max(ans, leftDepth + rightDepth)
		return max(leftDepth, rightDepth) + 1
	}
	getDepth(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
