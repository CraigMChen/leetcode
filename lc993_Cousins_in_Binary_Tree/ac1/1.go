package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 遍历到节点node时判断左右子树的值是否为两个目标值，如果是，则肯定不是堂兄弟
func isCousins(root *TreeNode, x int, y int) bool {
	var (
		xDepth, yDepth int
		isBrothers bool
		dfs func(node *TreeNode, curDepth int)
	)
	dfs = func(node *TreeNode, curDepth int) {
		if node == nil {
			return
		}
		if isBrothers {
			return
		}
		if node.Left != nil && (node.Left.Val == x || node.Left.Val == y) &&
			node.Right != nil && (node.Right.Val == y || node.Right.Val == x) {
			isBrothers = true
			return
		}
		if node.Val == x {
			xDepth = curDepth
		}
		if node.Val == y {
			yDepth = curDepth
		}
		dfs(node.Left, curDepth + 1)
		dfs(node.Right, curDepth + 1)
	}
	dfs(root, 0)
	return !isBrothers && xDepth == yDepth
}
