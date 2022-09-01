package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
// pA 从 headA 开始，pB 从 headB 开始，二者每次同时走一步
// pA 为空时，走到 headB，pB 为空时，走到 headA
// pA 和 pB 同时为空则不相交，pA 等于 pB 时返回
// 证明：设 headA 到交点的长度为 a，headB 到交点的长度为 b，两链表相交的长度为 c
// pA 从 headA 走到交点，再走过相交的段，再从 headB 走到交点，走过的长度为 a+c+b
// pB 从 headB 走到交点，再走过相交的段，在从 headA 走到交点，走过的长度为 b+c+a
// 上面两式是相等的，即当他们第二次走到交点的时候一定会相遇
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
