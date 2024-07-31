package test

type ListNode struct {
	Val int
	Next *ListNode
}

// 合并且去重
func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := new(ListNode)
	l3 := res
	isRoot := true
	for l1 != nil && l2 != nil {
		tmp := new(ListNode)
		if l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
		} else {
			tmp = l2
			l2 = l2.Next
		}
		if tmp.Val != l3.Val || isRoot {
			isRoot = false
			l3.Next = tmp
			l3 = l3.Next
		}
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return res.Next
}
