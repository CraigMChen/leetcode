package ac1

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func preorder(root *Node) []int {
	var (
		ans []int
		pre func(root *Node)
	)
	pre = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			pre(child)
		}
	}
	pre(root)
	return ans
}
