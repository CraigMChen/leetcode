package test2

type ListNode struct {
	Val int
	Next *ListNode
}

// 合并且去重 递归版
// 不知道对不对
func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else if l1.Val > l2.Val {
		l2.Next = merge(l1, l2.Next)
		return l2
	} else {
		l1.Next = merge(l1.Next, l2.Next)
		return l1
	}
}
