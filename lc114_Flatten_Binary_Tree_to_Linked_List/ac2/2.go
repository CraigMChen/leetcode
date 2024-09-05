package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 找前驱节点
// 不需要递归，直接从根节点开始遍历，节省栈空间
func flatten(root *TreeNode) {
	for root != nil {
		pre := root.Left
		if pre == nil {
			root = root.Right
			continue
		}
		for pre.Right != nil {
			pre = pre.Right
		}
		right := root.Right
		root.Left, root.Right = nil, root.Left
		pre.Right = right
		root = root.Right
	}
}
