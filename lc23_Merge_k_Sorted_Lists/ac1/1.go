package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 每次从数组中选出最小的节点
// 时间复杂度 O(nk)
// 空间复杂度 O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	var (
		parent, res *ListNode
	)
	getMinPtr := func() (*ListNode, int) {
		var (
			minPtr *ListNode
			minId  int
		)
		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if minPtr == nil || lists[i].Val < minPtr.Val {
				minPtr = lists[i]
				minId = i
			}
		}
		return minPtr, minId
	}

	minPtr, minId := getMinPtr()
	for minPtr != nil {
		if parent == nil {
			parent = minPtr
			res = parent
		} else {
			parent.Next = minPtr
			parent = minPtr
		}
		lists[minId] = lists[minId].Next
		minPtr, minId = getMinPtr()
	}
	return res
}
