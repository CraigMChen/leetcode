package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转子链表后连接头尾
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 辅助节点
	dummyNode := &ListNode{Next: head}
	parent := dummyNode
	var l, r, leftNode, rightNode *ListNode
	node := head
	// 反转子链表
	for i := 1; i <= right; i++ {
		if i < left {
			parent = node
			node = node.Next
			continue
		}
		if i == left {
			l = parent
			leftNode = node
		}
		if i == right {
			r = node.Next
			rightNode = node
		}
		next := node.Next
		node.Next = parent
		parent = node
		node = next
	}
	// 连接头
	if l != nil {
		l.Next = rightNode
	}
	// 连接尾
	leftNode.Next = r
	return dummyNode.Next
}
