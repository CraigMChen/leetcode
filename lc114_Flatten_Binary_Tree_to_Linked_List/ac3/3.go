package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 将左子树作为当前节点的右儿子
// 将右子树作为当前节点原左子树中最右侧叶子节点的右儿子
// 返回右子树中最右侧的叶子节点
func flatten(root *TreeNode) {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return node
		}
		left := dfs(node.Left)   // 左子树的最右侧节点
		right := dfs(node.Right) // 右子树的最右侧节点
		if left != nil {
			tmp := node.Right
			node.Right = node.Left
			node.Left = nil
			left.Right = tmp
		}
		if right == nil { // 若右子树的最右侧节点为空，则新树的最右侧节点为原左子树的最右侧节点
			return left
		}
		return right
	}
	dfs(root)
}
