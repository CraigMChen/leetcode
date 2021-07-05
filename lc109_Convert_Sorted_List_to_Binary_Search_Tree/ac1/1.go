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
