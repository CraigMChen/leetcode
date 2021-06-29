package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
// 题解：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/cong-qian-xu-yu-zhong-xu-bian-li-xu-lie-gou-zao-9/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	index := 0 // index表示中序遍历结果的下标
	stack := []*TreeNode{root} // 栈中的节点为 栈顶节点的所有未考虑过右儿子的祖先节点

	for i := 1; i < len(preorder); i++ { // 遍历前序遍历结果
		node := stack[len(stack) - 1]
		if node.Val != inorder[index] { // 若栈顶节点与当前中序遍历节点不相同
			left := &TreeNode{Val: preorder[i]}
			node.Left = left // 则当前遍历的节点一定是栈顶节点的左儿子
			stack = append(stack, left) // 将当前遍历的节点入栈
		} else { // 若栈顶节点与当前中序遍历节点相同
			for len(stack) != 0 && stack[len(stack) - 1].Val == inorder[index] { // 不断pop，直到栈为空或栈顶节点不等于当前中序遍历节点
				node = stack[len(stack) - 1] // 保存栈顶节点
				stack = stack[:len(stack) - 1] // 将栈顶节点出栈
				index++ // 遍历下一个中序遍历节点
			}
			right := &TreeNode{Val: preorder[i]}
			node.Right = right // 当前遍历的节点一定是最后一个出栈的节点的右儿子
			stack = append(stack, right) // 将当前遍历节点入栈
		}
	}
	return root
}
