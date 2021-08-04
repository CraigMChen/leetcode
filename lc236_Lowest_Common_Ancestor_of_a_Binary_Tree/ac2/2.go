package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 该函数可以理解为，找出以root为根的子树中，p和q的最近公共祖先，若没有最近公共祖先，则返回p或q的位置；若没有找到p或q，则返回空
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q { // 返回p或q的位置
		return root
	}
	// 递归找左右子树
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil { // 如果左右子树都有p或q，则当前节点一定是p和q的最近公共祖先
		return root
	}
	if l == nil { // 如果左子树的结果为空，则右子树的结果可能是p或q，也可能是p和q的最近公共祖先
		return r
	}
	return l
}
