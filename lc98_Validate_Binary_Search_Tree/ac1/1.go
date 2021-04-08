package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var isValid func(root, max, min *TreeNode) bool
	isValid = func(root, max, min *TreeNode) bool {
		if root == nil {
			return true
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		return isValid(root.Left, root, min) && isValid(root.Right, max, root)
	}
	return isValid(root, nil, nil)
}
