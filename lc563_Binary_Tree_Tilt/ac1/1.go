package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 递归求出各子树中所有节点的和，分别将左右子树的差的绝对值累加
func findTilt(root *TreeNode) int {
	var (
		ans int
		sum func(node *TreeNode) int
	)
	sum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return node.Val
		}
		left := sum(node.Left)
		right := sum(node.Right)
		ans += abs(left - right)
		return left + right + node.Val
	}
	sum(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
