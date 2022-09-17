package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 小顶堆
// 用 lists 建立小顶堆
// 然后用根节点建链表
// 将根节点指针指向它的 Next 并整理小顶堆，如果 Next 为空就把最后一个元素换上来，长度减一
// 不断重复直到堆中没有元素
// 时间复杂度 O(nlogk)
// 空间复杂度 O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	var notEmpty []*ListNode
	for _, node := range lists {
		if node != nil {
			notEmpty = append(notEmpty, node)
		}
	}
	lists = notEmpty

	var minHeapify func(int, int)
	minHeapify = func(size, i int) {
		l, r, min := i<<1+1, i<<1+2, i
		if l < size && lists[l].Val < lists[min].Val {
			min = l
		}
		if r < size && lists[r].Val < lists[min].Val {
			min = r
		}
		if min != i {
			lists[i], lists[min] = lists[min], lists[i]
			minHeapify(size, min)
		}
	}
	for i := len(lists)>>1 - 1; i >= 0; i-- {
		minHeapify(len(lists), i)
	}

	var parent, res *ListNode
	for len(lists) > 0 {
		if parent == nil {
			parent = lists[0]
			res = parent
		} else {
			parent.Next = lists[0]
			parent = lists[0]
		}
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists[0], lists[len(lists)-1] = lists[len(lists)-1], lists[0]
			lists = lists[:len(lists)-1]
		}
		minHeapify(len(lists), 0)
	}
	return res
}
