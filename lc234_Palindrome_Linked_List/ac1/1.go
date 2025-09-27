package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针 + 反转链表
// 先用快慢指针求出链表的中点
// 翻转从中点到结尾的链表
// 分别从起点和原链表最后一个节点开始比较是否相同
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	reverse(slow, fast)
	for node1, node2 := head, fast; node1 != nil && node2 != nil; node1, node2 = node1.Next, node2.Next {
		if node1.Val != node2.Val {
			return false
		}
	}
	return true
}

func reverse(start, end *ListNode) {
	var parent *ListNode
	last := end.Next
	for node := start; node != last; {
		next := node.Next
		node.Next = parent
		parent = node
		node = next
	}
}
