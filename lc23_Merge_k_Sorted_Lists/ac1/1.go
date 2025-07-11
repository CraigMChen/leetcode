package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 顺序合并
// 按照顺序两两合并
// 时间复杂度 O(n*k^2)
// 空间复杂度 O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for _, list := range lists {
		res = mergeLists(res, list)
	}
	return res
}

func mergeLists(root1 *ListNode, root2 *ListNode) *ListNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	dummy := &ListNode{}
	parent := dummy
	for root1 != nil && root2 != nil {
		if root1.Val <= root2.Val {
			parent.Next = root1
			root1 = root1.Next
		} else {
			parent.Next = root2
			root2 = root2.Next
		}
		parent = parent.Next
	}
	if root1 != nil {
		parent.Next = root1
	}
	if root2 != nil {
		parent.Next = root2
	}
	return dummy.Next
}
