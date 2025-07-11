package ac2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 小顶堆（优先队列）
// 用 lists 建立小顶堆
// 每次将堆顶元素的节点插入新链表
// 然后将堆顶指针指针指向它的 Next 并整理小顶堆，如果 Next 为空就把最后一个元素换上来，长度减一
// 不断重复直到所有节点都插入新链表
// 时间复杂度 O(nklogk)
// 空间复杂度 O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	heap := make([]*ListNode, 0, len(lists))
	for _, list := range lists {
		if list != nil {
			heap = append(heap, list)
		}
	}

	for i := (len(heap) - 2) >> 1; i >= 0; i-- {
		heapify(heap, i)
	}

	dummy := &ListNode{}
	parent := dummy
	for len(heap) > 0 {
		parent.Next = heap[0]
		heap[0] = heap[0].Next
		if heap[0] == nil {
			heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
			heap = heap[:len(heap)-1]
		}
		heapify(heap, 0)
		parent = parent.Next
	}
	return dummy.Next
}

func heapify(lists []*ListNode, i int) {
	for i<<1+1 < len(lists) {
		l, r, m := i<<1+1, i<<1+2, i
		if l < len(lists) && lists[l].Val < lists[m].Val {
			m = l
		}
		if r < len(lists) && lists[r].Val < lists[m].Val {
			m = r
		}
		if m == i {
			return
		}
		lists[i], lists[m] = lists[m], lists[i]
		i = m
	}
}
