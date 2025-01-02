package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{Next: list1}
	pre := dummyNode
	node1, node2 := list1, list2
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			pre = node1
			node1 = node1.Next
		} else {
			pre.Next = node2
			next := node2.Next
			node2.Next = node1
			pre = node2
			node2 = next
		}
	}
	if node2 != nil {
		pre.Next = node2
	}
	return dummyNode.Next
}
