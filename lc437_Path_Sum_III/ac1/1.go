package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回溯
func pathSum(root *TreeNode, targetSum int) int {
	var (
		ans int
		dfs func(node *TreeNode, cur int)
	)
	count := map[int]int{0: 1} // 前缀和的出现次数

	dfs = func(node *TreeNode, cur int) { // cur表示当前的前缀和
		if node == nil {
			return
		}
		cur += node.Val
		ans += count[cur-targetSum] // 从之前的前缀和中找到cur-targetSum，如果存在则加上该前缀和出现次数
		count[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		count[cur]-- // 回溯
	}
	dfs(root, 0)
	return ans
}
