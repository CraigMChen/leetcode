package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针+反转链表+合并链表
// 首先用快慢指针求出链表的中点
// 然后将中点之前的节点的 Next 指针置空，并将中点开始到结尾的链表翻转
// 将翻转后的链表与原链表的前半部分合并
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 快慢指针求中点
	var (
		fast, slow = head, head
		parent     *ListNode
	)
	for fast != nil {
		parent = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	// 将中点前面节点的 Next 指针置空
	parent.Next = nil

	// 翻转链表
	parent = nil
	for slow != nil {
		slow, parent, slow.Next = slow.Next, slow, parent
	}

	// 合并链表
	for p1, p2 := head, parent; p2 != nil; {
		p1, p2, p1.Next, p2.Next = p1.Next, p2.Next, p2, p1.Next
	}
}
