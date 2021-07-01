package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
// 方法与lc105 ac2相同
func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(inorder)
	if length == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[length - 1]}
	stack := []*TreeNode{root}
	index := length - 1
	for i := length - 2; i >= 0; i-- {
		node := stack[len(stack) - 1]
		if node.Val != inorder[index] {
			right := &TreeNode{Val: postorder[i]}
			node.Right = right
			stack = append(stack, right)
		} else {
			for len(stack) != 0 && stack[len(stack) - 1].Val == inorder[index] {
				node = stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				index--
			}
			left := &TreeNode{Val: postorder[i]}
			node.Left = left
			stack = append(stack, left)
		}
	}
	return root
}
