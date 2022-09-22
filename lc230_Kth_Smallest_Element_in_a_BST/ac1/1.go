package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func kthSmallest(root *TreeNode, k int) int {
	var inOrder func(node *TreeNode)
	n := 0
	ans := 0
	inOrder = func(node *TreeNode) {
		if n >= k || node == nil { // 剪枝
			return
		}
		inOrder(node.Left)
		n++
		if n == k {
			ans = node.Val
			return
		}
		inOrder(node.Right)
	}
	inOrder(root)
	return ans
}
