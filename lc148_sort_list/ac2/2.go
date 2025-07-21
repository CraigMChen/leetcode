package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 自顶向下归并排序
// 用快慢指针找到链表的终点
// 分别对左右两边的链表进行排序，然后合并
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var parent *ListNode
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		parent = slow
		slow = slow.Next
	}
	if parent != nil {
		parent.Next = nil
	}
	left := sortList(head)
	right := sortList(slow)
	return mergeList(left, right)
}

func mergeList(node1, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	parent := dummy
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			parent.Next = node1
			node1 = node1.Next
		} else {
			parent.Next = node2
			node2 = node2.Next
		}
		parent = parent.Next
	}
	if node1 != nil {
		parent.Next = node1
	}
	if node2 != nil {
		parent.Next = node2
	}
	return dummy.Next
}
