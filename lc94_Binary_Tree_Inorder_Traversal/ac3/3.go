package ac3

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Morris
func inorderTraversal(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			suc := root.Left
			for suc.Right != nil && suc.Right != root {
				suc = suc.Right
			}
			if suc.Right == nil {
				suc.Right = root
				root = root.Left
			} else {
				suc.Right = nil
				res = append(res, root.Val)
				root = root.Right
			}
		}
	}
	return res
}
