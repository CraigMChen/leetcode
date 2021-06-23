package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 一颗树为BST的条件：
// 左子树为BST
// 右子树为BST
// 左子树中所有节点的值都小于根节点
// 右子树中所有节点的值都大于根节点
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
