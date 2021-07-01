package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	length := len(inorder)
	root := &TreeNode{Val: postorder[length - 1]}
	index := -1
	for i := range inorder {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index + 1 : length], postorder[index : length - 1])
	return root
}
