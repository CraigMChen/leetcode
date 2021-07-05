package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归/分治
// 链表的中间节点一定是二叉搜索树的根节点
// 找中键节点可以使用快慢指针的方法：
// 从头指针开始，快指针走两步，慢指针走一步
// 当快指针走到最后一个节点时，慢指针所指的节点就是中间节点
func sortedListToBST(head *ListNode) *TreeNode {
	var buildBST func(left, right *ListNode) *TreeNode
	buildBST = func(left, right *ListNode) *TreeNode {
		if left == right {
			return nil
		}
		low, fast := left, left
		for fast.Next != right && fast.Next != nil {
			fast = fast.Next
			if fast.Next != right && fast.Next != nil {
				fast = fast.Next
			}
			low = low.Next
		}

		root := &TreeNode{Val: low.Val}
		root.Left = buildBST(left, low)
		root.Right = buildBST(low.Next, right)
		return root
	}
	return buildBST(head, nil)
}
