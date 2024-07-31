package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力 优化
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	res := l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	} else if l2 != nil {
		l.Next = l2
	}
	return res.Next
}

// 优化2
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	parent := &ListNode{Next: list1}
//	root := parent
//	for list1 != nil || list2 != nil {
//		if list1 == nil {
//			parent.Next = list2
//			list2 = list2.Next
//		} else if list2 == nil {
//			parent.Next = list1
//			list1 = list1.Next
//		} else if list1.Val <= list2.Val {
//			parent.Next = list1
//			list1 = list1.Next
//		} else {
//			parent.Next = list2
//			list2 = list2.Next
//		}
//		parent = parent.Next
//	}
//	return root.Next
//}
