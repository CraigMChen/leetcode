package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
// BST的中序遍历结果是非递减序列，最相近的两个值一定相邻
func getMinimumDifference(root *TreeNode) int {
	ans := 1 << 31 - 1
	var (
		inorder func(root *TreeNode)
		pre *TreeNode = nil
	)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if pre != nil {
			diff := abs(root.Val - pre.Val)
			if diff < ans {
				ans = diff
			}
		}
		pre = root
		inorder(root.Right)
	}
	inorder(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
