package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
// pA 从 headA 开始，pB 从 headB 开始，二者每次同时走一步
// pA 为空时，走到 headB，pB 为空时，走到 headA
// pA 和 pB 同时为空则不相交，pA 等于 pB 时返回
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
