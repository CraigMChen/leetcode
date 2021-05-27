package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从根节点往下遍历，如果p和q分别在左右子树，则当前节点为答案
// 如果p和q都在右子树，则向右子树遍历
// 否则向左子树遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val <= root.Val && q.Val >= root.Val ||
			p.Val >= root.Val && q.Val <= root.Val {
			return root
		}
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
