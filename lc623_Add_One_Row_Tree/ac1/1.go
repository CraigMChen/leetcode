package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 深搜
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	var dfs func(node *TreeNode, cur int)
	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		if cur == depth-1 {
			right := node.Right
			left := node.Left
			node.Right = &TreeNode{
				Val:   val,
				Right: right,
			}
			node.Left = &TreeNode{
				Val:  val,
				Left: left,
			}
			return
		}
		dfs(node.Right, cur+1)
		dfs(node.Left, cur+1)
	}
	dfs(root, 1)
	return root
}
