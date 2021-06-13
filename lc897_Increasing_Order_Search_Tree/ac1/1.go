package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构造新树
func increasingBST(root *TreeNode) *TreeNode {
	var (
		ans, pre *TreeNode
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if ans == nil {
			ans, pre = node, node
		} else {
			pre.Right = &TreeNode{
				Val:   node.Val,
			}
			pre = pre.Right
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
