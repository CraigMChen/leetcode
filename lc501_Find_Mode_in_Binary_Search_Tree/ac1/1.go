package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 时间O(n)
// 空间O(n)
// BST的中序遍历是一个非递减序列，相同的数字肯定在一起
func findMode(root *TreeNode) []int {
	var (
		ans []int
		inorder func(root *TreeNode)
		update func(x int)
	)
	cur, count, maxCount := -1, 0, 0
	inorder = func(root *TreeNode) { // 对BST进行中序遍历
		if root == nil {
			return
		}
		inorder(root.Left)
		update(root.Val) // 遍历时统计每个数的出现次数
		inorder(root.Right)
	}
	update = func(x int) {
		if x != cur { // 如果当前遍历的值与上一个值不同
			count = 1 // 则将count置为1
			cur = x
		} else {
			count++ // 否则count加一
		}
		if count > maxCount { // 如果当前数字的出现次数大于出现最多的次数
			ans = []int{cur}  // 则当前遍历的值为新的众数
			maxCount = count  // 更新出现最多的次数
		} else if count == maxCount { // 如果当前出现次数和最大次数相同，则该数也是众数
			ans = append(ans, cur)
		}
	}
	inorder(root)
	return ans
}
