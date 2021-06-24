package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// 交换顺序后，中序遍历会有一处（交换相邻两个数）或两处出现ai>ai+1的情况
// 用pre维护前驱节点
// 判断当前节点是否小于前驱节点
// 用x和y将异常的节点记录下来，最后交换即可
func recoverTree(root *TreeNode)  {
	var (
		pre, x, y *TreeNode
		inorder func(node *TreeNode)
	)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if pre != nil && node.Val < pre.Val {
			y = node
			if x == nil {
				x = pre
			} else {
				return
			}
		}
		pre = node
		inorder(node.Right)
	}
	inorder(root)
	x.Val, y.Val = y.Val, x.Val
}
