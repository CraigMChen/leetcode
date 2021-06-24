package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris中序遍历
// 用Morris时不能提前break，因为修改了原树，必须复原
// 否则判题代码会进入无限循环
func recoverTree(root *TreeNode) {
	var (
		pre, x, y *TreeNode
		check func(cur *TreeNode)
	)
	check = func(cur *TreeNode) {
		if pre != nil && cur.Val < pre.Val {
			y = cur
			if x == nil {
				x = pre
			}
		}
		pre = cur
	}

	for root != nil {
		if root.Left == nil {
			check(root)
			root = root.Right
			continue
		}

		predecessor := root.Left
		for predecessor.Right != nil && predecessor.Right != root {
			predecessor = predecessor.Right
		}
		if predecessor.Right == nil {
			predecessor.Right = root
			root = root.Left
		} else {
			predecessor.Right = nil
			check(root)
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}
