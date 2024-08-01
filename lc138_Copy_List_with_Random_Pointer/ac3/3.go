package ac3

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 迭代
// 将每个原节点复制为两个相邻的节点，如原链表为 A -> B -> C
// 复制后的链表为 A -> A' -> B -> B' -> C -> C'
// 其中 X' 为 X 的复制节点
// 新的 Next 指针起到了前面方法 map 的效果，即将旧节点与新节点映射起来了
// 接下来只要按照方法 1 中的步骤进行操作即可
// 注意：需要将旧链表还原，否则影响判题逻辑
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 复制每个节点并按顺序连接
	for node := head; node != nil; {
		newNode := &Node{
			Val:  node.Val,
			Next: node.Next,
		}
		node.Next = newNode
		node = newNode.Next
	}
	res := head.Next // 新的链表头即为旧的链表头的复制节点
	// 为避免前面的 Next 指针复原后影响到后面的 Random 指针
	// 先处理 Random 指针
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	// 最后处理 Next 指针，并将旧链表的 Next 指针复原
	for node := head; node != nil; {
		next := node.Next.Next
		if next != nil {
			node.Next.Next = next.Next
		}
		node.Next = next
		node = next
	}
	return res
}
