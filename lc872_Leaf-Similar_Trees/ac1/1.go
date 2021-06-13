package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := getLeaves(root1)
	leaves2 := getLeaves(root2)
	return isEqual(leaves1, leaves2)
}

func getLeaves(root *TreeNode) []int {
	var (
		res []int
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
