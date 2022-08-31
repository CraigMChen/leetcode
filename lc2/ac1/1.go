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
