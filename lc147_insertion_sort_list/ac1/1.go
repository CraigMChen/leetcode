package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	parent := dummy
	for cur := head; cur != nil; {
		node := dummy.Next
		last := dummy
		for node != cur {
			if cur.Val < node.Val {
				break
			}
			last = node
			node = node.Next
		}
		if node == cur {
			parent = cur
			cur = cur.Next
			continue
		}
		next := cur.Next
		parent.Next = next
		last.Next = cur
		cur.Next = node
		cur = next
	}
	return dummy.Next
}
