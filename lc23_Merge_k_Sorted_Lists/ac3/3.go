package ac3

type ListNode struct {
	Val  int
	Next *ListNode
}

// 分治
// 时间复杂度 O(kn*log n)
// 空间复杂度 O(log n)
func mergeKLists(lists []*ListNode) *ListNode {
	merge2Lists := func(a, b *ListNode) *ListNode {
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		var parent, res *ListNode
		for a != nil && b != nil {
			if a.Val < b.Val {
				if parent == nil {
					parent = a
					res = a
				} else {
					parent.Next = a
					parent = a
				}
				a = a.Next
			} else {
				if parent == nil {
					parent = b
					res = b
				} else {
					parent.Next = b
					parent = b
				}
				b = b.Next
			}
		}
		if a != nil {
			parent.Next = a
		} else {
			parent.Next = b
		}
		return res
	}

	var merge func(int, int) *ListNode
	merge = func(l, r int) *ListNode {
		if l == r-1 {
			return lists[l]
		}
		if l >= r {
			return nil
		}
		m := (r-l)>>1 + l
		return merge2Lists(merge(l, m), merge(m, r))
	}
	return merge(0, len(lists))
}
