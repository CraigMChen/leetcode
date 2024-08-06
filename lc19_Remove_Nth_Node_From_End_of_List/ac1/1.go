package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// 快指针先走 n-1 步
// 然后快慢指针同时走，每次走 1 步，直到快指针走到最后一个节点时
// 慢指针指向的节点即为倒数第 n 个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	parent := dummyNode
	slow, fast := head, head
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		parent = slow
		slow = slow.Next
	}
	parent.Next = slow.Next
	return dummyNode.Next
}
