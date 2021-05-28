package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜，用一个bool标记是否为左子节点
func sumOfLeftLeaves(root *TreeNode) int {
	var (
		sum int
		getSum func(node *TreeNode, isLeft bool)
	)

	getSum = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			sum += node.Val
			return
		}
		getSum(node.Left, true)
		getSum(node.Right, false)
	}
	getSum(root, false)
	return sum
}
