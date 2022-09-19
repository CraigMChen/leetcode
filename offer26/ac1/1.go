package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 首先遍历 A 树，找到一个值与 B 的根节点值相等的节点
// 然后判断该节点是否与 B 树的值相等
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	var isEqual func(a, b *TreeNode) bool
	isEqual = func(a, b *TreeNode) bool {
		if a == nil && b != nil {
			return false
		}
		if b == nil {
			return true
		}
		if a.Val != b.Val {
			return false
		}
		return isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
	}

	var isSub func(a *TreeNode) bool
	isSub = func(a *TreeNode) bool {
		if a == nil {
			return false
		}
		if a.Val == B.Val && isEqual(a, B) {
			return true
		}
		return isSub(a.Left) || isSub(a.Right)
	}
	return isSub(A)
}
