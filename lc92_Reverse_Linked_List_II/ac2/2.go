package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 头插法
// 1 -> 2 -> 3 -> 4 -> 5 -> 6  (2,5)
// 首先将 3 插到 2 的前面
// 1 -> 3 -> 2 -> 4 -> 5 -> 6
// 然后将 4 插到 3 的前面
// 1 -> 4 -> 3 -> 2 -> 5 -> 6
// 最后将 5 插到 4 的前面
// 1 -> 5 -> 4 -> 3 -> 2 -> 6
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	node := head
	pre := dummyNode
	var cur, leftNode *ListNode
	for i := 1; i <= right; i++ {
		if i < left {
			pre = node
			node = node.Next
			continue
		}
		if i == left {
			leftNode = node
			cur = node
			node = node.Next
			continue
		}
		next := node.Next
		pre.Next = node
		node.Next = cur
		leftNode.Next = next
		cur = node
		node = next
	}
	return dummyNode.Next
}
