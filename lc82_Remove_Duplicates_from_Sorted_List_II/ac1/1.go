package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		p      = head
		parent *ListNode
	)
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			n := p
			for n != nil && n.Val == p.Val {
				n = n.Next
			}
			if parent == nil {
				head = n
			} else {
				parent.Next = n
			}
			p = n
		} else {
			parent = p
			p = p.Next
		}
	}
	return head
}
