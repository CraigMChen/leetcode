package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 时间O(n)
// 空间O(H)
func findMode(root *TreeNode) []int {
	var (
		ans []int
		inorder func(root *TreeNode)
		update func(x int)
	)
	cur, count, maxCount := -1, 0, 0
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		update(root.Val)
		inorder(root.Right)
	}
	update = func(x int) {
		if x != cur {
			count = 1
			cur = x
		} else {
			count++
		}
		if count > maxCount {
			ans = []int{cur}
			maxCount = count
		} else if count == maxCount {
			ans = append(ans, cur)
		}
	}
	inorder(root)
	return ans
}
