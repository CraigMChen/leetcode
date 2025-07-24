package ac3

type ListNode struct {
	Val  int
	Next *ListNode
}

// 自底向上的归并排序
// 对长度从 1 开始的链表，从左到右进行归并排序
// 每一轮过后长度翻倍，直到长度大于等于总长度
func sortList(head *ListNode) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	for l := 1; l < length; l <<= 1 {
		parent := dummy
		for node1l := dummy.Next; node1l != nil; {
			node1r := node1l
			node2l := node1l
			for i := 0; i < l && node2l != nil; i++ {
				node1r = node2l
				node2l = node2l.Next
			}
			if node2l == nil {
				break
			}
			node1r.Next = nil
			node2r := node2l
			for i := 1; i < l && node2r.Next != nil; i++ {
				node2r = node2r.Next
			}
			next := node2r.Next
			node2r.Next = nil
			var last *ListNode
			parent.Next, last = mergeList(node1l, node2l)
			last.Next = next
			parent = last
			node1l = next
		}
	}
	return dummy.Next
}

func mergeList(node1, node2 *ListNode) (*ListNode, *ListNode) {
	dummy := &ListNode{}
	parent := dummy
	var last *ListNode
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			parent.Next = node1
			node1 = node1.Next
			last = node1
		} else {
			parent.Next = node2
			node2 = node2.Next
			last = node2
		}
		parent = parent.Next
	}

	if node1 != nil {
		parent.Next = node1
		for ; node1 != nil; node1 = node1.Next {
			last = node1
		}
	}
	if node2 != nil {
		parent.Next = node2
		for ; node2 != nil; node2 = node2.Next {
			last = node2
		}
	}
	return dummy.Next, last
}
