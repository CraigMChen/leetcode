package ac3

type ListNode struct {
	Val  int
	Next *ListNode
}

// 分治
// 时间复杂度 O(nk*logk)
// 空间复杂度 O(logk)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	m := len(lists) >> 1
	l := mergeKLists(lists[:m])
	r := mergeKLists(lists[m:])
	return mergeLists(l, r)
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
