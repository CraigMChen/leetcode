package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 分治+中序遍历
// 有序链表的值的顺序即为二叉搜索树中序遍历的顺序
// 先遍历链表，求出链表的长度
// 再按照中序遍历的顺序：构建左子树，构建根节点，构建右子树
func sortedListToBST(head *ListNode) *TreeNode {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	cur = head
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left >= right {
			return nil
		}

		mid := (right - left) >> 1 + left
		root := &TreeNode{}
		root.Left = build(left, mid)
		root.Val = cur.Val
		cur = cur.Next
		root.Right = build(mid + 1, right)
		return root
	}
	return build(0, length)
}
