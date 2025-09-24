package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自底向上的递归
// 同时求节点的高度，以及以当前节点为根节点的子树是否为平衡二叉树
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func isBalanced(root *TreeNode) bool {
	_, ok := dfs(root)
	return ok
}

func dfs(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	l, lOk := dfs(node.Left)
	r, rOk := dfs(node.Right)
	return max(l, r) + 1, l-r >= -1 && l-r <= 1 && lOk && rOk
}
