package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将 Val 大于等于 x 的节点摘出，连接成新的链表
// 最后将新链表连接到旧链表的尾部
func partition(head *ListNode, x int) *ListNode {
	dummyNode := &ListNode{Next: head}
	bigHead := &ListNode{}
	bigNode := bigHead
	parent := dummyNode
	node := head
	for node != nil {
		next := node.Next
		if node.Val >= x {
			parent.Next = next
			bigNode.Next = node
			bigNode = node
			// 注意，需要将 Next 置空
			// 否则新链表的最后一个节点的 Next 指针还连接着旧链表的节点
			// 会导致链表成环
			bigNode.Next = nil
		} else {
			parent = node
		}
		node = next
	}
	parent.Next = bigHead.Next
	return dummyNode.Next
}
