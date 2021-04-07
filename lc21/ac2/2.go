package ac2

type ListNode struct {
	Val int
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
