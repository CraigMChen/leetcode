package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 分治
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var (
		build func(l, r int) *TreeNode
		max func(l, r int) (int, int)
	)
	max = func(l, r int) (int, int) {
		m := -1
		id := -1
		for i := l; i < r; i++ {
			if nums[i] > m {
				m = nums[i]
				id = i
			}
		}
		return m, id
	}

	build = func(l, r int) *TreeNode {
		if l >= r {
			return nil
		}
		m, id := max(l, r)
		root := &TreeNode{
			Val:   m,
			Left:  build(l, id),
			Right: build(id+1, r),
		}
		return root
	}
	return build(0, len(nums))
}
