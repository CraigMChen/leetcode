package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// fast 和 slow 指针分别从 head 开始走，slow 每次走一步，fast 每次走两步
// 设入环点的下标为 a，相遇点的下标为 b，环的长度为 c
// 相遇时，slow 走过的长度为 b
// fast 走过的长度为 b + c
// 在任意时刻，fast 走过的长度是 slow 的两倍，得 b+c=2b，即 b = c
// 即环的长度与相遇点距离起点的长度相同
// 从入环点出发，走 n 个完整的环，即 a + nc，一定会回到入环点
// a + nc = a + nb
// 相遇点已经在 b 上，从相遇点出发走长度 a 一定会到达入环点
// 从 head 出发走长度 a 也会到达入环点
// 那么两个指针分别从 head 和 b 开始走，他们相遇的地方就是入环点
func detectCycle(head *ListNode) *ListNode {
	fast, slow, ans := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	for ans != slow {
		ans = ans.Next
		slow = slow.Next
	}
	return ans
}
