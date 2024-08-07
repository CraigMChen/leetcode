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

// 2024.08.07
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	dummyNode := &ListNode{Next: head}
//	parent := dummyNode
//	for node := head; node != nil; {
//		var l, r *ListNode
//		l, r, node = reverse(node, k)
//		parent.Next = l
//		parent = r
//	}
//	return dummyNode.Next
//}
//
//func reverse(node *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
//	start := node
//	i := 1
//	for i < k && node.Next != nil {
//		node = node.Next
//		i++
//	}
//	end := node
//	next := end.Next
//	if i < k {
//		return start, end, nil
//	}
//
//	var parent *ListNode
//	for node = start; node != next; {
//		tmp := node.Next
//		node.Next = parent
//		parent = node
//		node = tmp
//	}
//	return end, start, next
//}
