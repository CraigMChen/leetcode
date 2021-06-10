package ac2

type Node struct {
	Val      int
	Children []*Node
}

// 迭代
// 用切片模拟栈
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ans []int
	stack := []*Node{root}
	for len(stack) != 0 {
		l := len(stack)
		node := stack[l - 1]
		stack = stack[:l - 1]
		ans = append(ans, node.Val)

		cl := len(node.Children)
		for i := cl - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ans
}
