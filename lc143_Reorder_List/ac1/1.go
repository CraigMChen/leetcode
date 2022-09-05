package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 线性表
// 用切片按序将每个节点保存下来
func reorderList(head *ListNode) {
	var nodes []*ListNode
	for p := head; p != nil; p = p.Next {
		nodes = append(nodes, p)
	}

	n := len(nodes)
	for i := (n + 1) / 2; i < n; i++ {
		nodes[n-i-1].Next = nodes[i]
		nodes[i].Next = nodes[n-i]
		if i == (n+1)/2 {
			nodes[n-i].Next = nil
		}
	}
}
