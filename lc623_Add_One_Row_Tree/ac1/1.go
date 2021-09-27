package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	var dfs func(node, father *TreeNode, isRight bool, cur int)
	dfs = func(node, father *TreeNode, isRight bool, cur int) {
		if node == nil && father == nil {
			return
		}
		if cur == depth {
			if isRight {
				father.Right = &TreeNode{
					Val:   val,
					Right: node,
				}
			} else {
				father.Left = &TreeNode{
					Val:  val,
					Left: node,
				}
			}
			return
		}
		if node != nil {
			dfs(node.Left, node, false, cur+1)
			dfs(node.Right, node, true, cur+1)
		}
	}
	dfs(root, nil, false, 1)
	return root
}
