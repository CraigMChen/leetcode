package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 递归求每个节点的高度
// height = max(height(root.Left), height(root.Right))+1
// 遍历节点的同时更新答案
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}
	dfs(root)
	return ans
}
