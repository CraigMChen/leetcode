package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
