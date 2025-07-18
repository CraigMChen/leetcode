package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 先对右子节点深搜再对左子节点深搜，这样每层第一次遇到的节点一定是改层右视图的节点
// 标记当前深度 level，以及答案数组
// 若 level 大于 ans 数组的长度，说明当前 level 是第一次遍历到，当前节点加入答案数组
func rightSideView(root *TreeNode) []int {
	var (
		ans []int
		dfs func(node *TreeNode, level int)
	)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level >= len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root, 0)
	return ans
}
