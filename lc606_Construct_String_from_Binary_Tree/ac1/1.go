package ac1

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// "省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对"的意思是
// 输出的答案只能构造唯一的二叉树，没有歧义
// 若左右子树都不为空，则都要加括号
// 若左子树不为空，右子树为空，则右子树不用加括号
// 若左子树为空，右子树不为空，则左子树需要加括号
func tree2str(root *TreeNode) string {
	var (
		ans string
		preorder func(node *TreeNode)
	)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans += strconv.FormatInt(int64(node.Val), 10)
		if node.Left != nil && node.Right != nil {
			ans += "("
			preorder(node.Left)
			ans += ")("
			preorder(node.Right)
			ans += ")"
		} else if node.Left != nil && node.Right == nil {
			ans += "("
			preorder(node.Left)
			ans += ")"
		} else if node.Left == nil && node.Right != nil {
			ans += "()("
			preorder(node.Right)
			ans += ")"
		}
	}
	preorder(root)
	return ans
}
