package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 前序遍历结果的第一个一定是根节点
// 从中序遍历中找到根节点的位置就可以确定左右子树的节点数量
// 从而分别确定左右子树的前序遍历结果和中序遍历结果
// 把问题分为两个子问题，不断递归下去
// 递归终点是中序遍历或前序遍历结果为空
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}
	headIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			headIndex = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:headIndex+1], inorder[:headIndex]),
		Right: buildTree(preorder[headIndex+1:], inorder[headIndex+1:]),
	}
}
