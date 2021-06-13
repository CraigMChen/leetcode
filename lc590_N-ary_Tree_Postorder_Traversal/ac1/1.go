package ac1

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func postorder(root *Node) []int {
	var (
		ans []int
		post func(node *Node)
	)
	post = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			post(child)
		}
		ans = append(ans, node.Val)
	}
	post(root)
	return ans
}
