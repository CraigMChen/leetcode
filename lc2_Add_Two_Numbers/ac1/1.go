package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res    = l1
		parent *ListNode
		carry  = 0
	)
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			l1 = new(ListNode)
			parent.Next = l1
		}
		add := 0
		if l2 != nil {
			add = l2.Val
		}
		sum := l1.Val + add + carry
		carry, l1.Val = sum/10, sum%10

		parent = l1
		l1 = l1.Next
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return res
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	carry := 0
//	var parent *ListNode
//	ans := l1
//	for l1 != nil || l2 != nil {
//		a, b := 0, 0
//		if l1 == nil {
//			l1 = &ListNode{}
//			parent.Next = l1
//		}
//		a = l1.Val
//		if l2 != nil {
//			b = l2.Val
//		}
//		l1.Val = (a + b + carry) % 10
//		carry = (a + b + carry) / 10
//		parent = l1
//		l1 = l1.Next
//		if l2 != nil {
//			l2 = l2.Next
//		}
//	}
//	if carry > 0 {
//		parent.Next = &ListNode{Val: carry}
//	}
//	return ans
//}
