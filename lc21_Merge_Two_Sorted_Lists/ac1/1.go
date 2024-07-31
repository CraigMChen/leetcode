package ac1

type ListNode struct {
	Val int
	Next *ListNode
}

// 暴力
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l := new(ListNode)
	res := l
	var father *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Val = l1.Val
			l1 = l1.Next
		} else {
			l.Val = l2.Val
			l2 = l2.Next
		}
		father = l
		l.Next = new(ListNode)
		l = l.Next
	}
	for l1 != nil {
		l.Val = l1.Val
		l1 = l1.Next
		father = l
		l.Next = new(ListNode)
		l = l.Next
	}
	for l2 != nil {
		l.Val = l2.Val
		l2 = l2.Next
		father = l
		l.Next = new(ListNode)
		l = l.Next
	}
	father.Next = nil
	return res
}
