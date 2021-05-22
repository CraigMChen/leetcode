package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func sortedArrayToBST(nums []int) *TreeNode {
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left >= right {
			return nil
		}
		mid := (right - left) >> 1 + left
		root := &TreeNode{Val: nums[mid]}
		root.Left = build(left, mid)
		root.Right = build(mid + 1, right)
		return root
	}
	return build(0, len(nums))
}
