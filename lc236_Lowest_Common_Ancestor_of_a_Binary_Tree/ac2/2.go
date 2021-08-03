package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q { // 如果当前节点为p或q，则返回当前节点
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil { // 如果左右子树都包含p或q，则返回当前节点
		return root
	}
	if l == nil {
		return r
	}
	return l
}
