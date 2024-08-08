package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	parent := dummyNode
	for node := head; node != nil; {
		if node.Next == nil || node.Val != node.Next.Val {
			parent = node
			node = node.Next
			continue
		}
		for node.Next != nil && node.Val == node.Next.Val {
			node = node.Next
		}
		parent.Next = node.Next
		node = node.Next
	}
	return dummyNode.Next
}
