package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	father := l1
	ans := l1

	for l1 != nil || l2 != nil {
		if l2 == nil && carry == 0 {
			break
		}
		if l1 == nil {
			l1 = new(ListNode)
			father.Next = l1
		}
		var sum int
		if l2 != nil {
			sum = l1.Val + l2.Val + carry
		} else {
			sum = l1.Val + carry
		}
		l1.Val = sum % 10
		carry = sum / 10
		father = l1
		l1 = l1.Next
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		l1 = &ListNode{carry, nil}
		father.Next = l1
	}
	return ans
}
