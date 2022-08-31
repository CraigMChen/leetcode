package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// fast 指针每次移动 2 步，slow 指针每次移动 1 步
// 设 fast 指针和 slow 指针移动的节点数分别为 f 和 s，入环点距离 head 长度为 a，环的长度为 b
// 当 fast 指针和 slow 指针相遇时，有 f = s + nb
// fast 的速度是 slow 的 2 倍，有 f = 2s
// 根据上面两式，可以得出
// s = nb, f = 2nb，即 slow 走了 n 圈的环，fast 走了 2n 圈的环
// slow 指针从 head 走到入环点的距离 k = a + nb（走过距离 a 后再走 n 全都在入环点上）
// 而此时 slow 已经走了 nb，再走长度 a 就是入环点了，a 就是要求的距离
// 指针 ptr 从 head 开始和 slow 一起走，每次移动一步，他们相遇时 ptr 的位置就是入环点
func detectCycle(head *ListNode) *ListNode {
	fast, slow, ptr := head, head, head
	for {
		if slow == nil || fast == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if slow == fast {
			break
		}
	}
	for ptr != slow {
		ptr = ptr.Next
		slow = slow.Next
	}
	return ptr
}
