package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		p      = head
		parent = head
	)
	head = p.Next
	p.Next, p.Next.Next = p.Next.Next, p
	p = p.Next
	for p != nil {
		if p.Next == nil {
			break
		}
		parent.Next = p.Next
		p.Next, p.Next.Next = p.Next.Next, p
		parent = p
		p = p.Next
	}
	return head
}
