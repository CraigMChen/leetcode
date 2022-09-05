package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 模拟
func reverseKGroup(head *ListNode, k int) *ListNode {
	reverse := func(h *ListNode) (*ListNode, *ListNode) {
		var (
			i      = 0
			p      = h
			father *ListNode
		)
		for ; i < k && p != nil; i++ {
			p = p.Next
		}
		if i < k {
			return h, nil
		}

		p = h
		for i := 0; i < k; i++ {
			tmp := p.Next
			p.Next = father
			father = p
			p = tmp
		}
		return father, p
	}

	var (
		parent, h *ListNode
		p         = head
	)
	for p != nil {
		tmp := p
		h, p = reverse(p)
		if parent == nil {
			head = h
		} else {
			parent.Next = h
		}
		parent = tmp
	}
	return head
}
