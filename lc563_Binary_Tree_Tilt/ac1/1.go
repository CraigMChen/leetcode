package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
