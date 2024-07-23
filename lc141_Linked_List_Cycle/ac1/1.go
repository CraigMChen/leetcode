package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
	}
	return slow == fast
}
