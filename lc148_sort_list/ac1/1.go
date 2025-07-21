package ac1

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 辅助数组
// 将链表各节点放入辅助数组
// 对辅助数组排序后重新建链表
func sortList(head *ListNode) *ListNode {
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	dummy := &ListNode{}
	parent := dummy
	for i := 0; i < len(nodes); i++ {
		parent.Next = nodes[i]
		parent = nodes[i]
	}
	parent.Next = nil
	return dummy.Next
}
