package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 旋转 k 次等价于旋转 k % n 次（n为链表长度）
// 将链表后 k 个节点移动到头部
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	k %= length
	if k == 0 {
		return head
	}

	newHead, oldLast, newLast := head, head, (*ListNode)(nil)
	for i := 1; i < k; i++ {
		oldLast = oldLast.Next
	}
	for oldLast.Next != nil {
		newLast = newHead
		newHead = newHead.Next
		oldLast = oldLast.Next
	}
	oldLast.Next = head
	newLast.Next = nil
	return newHead
}
