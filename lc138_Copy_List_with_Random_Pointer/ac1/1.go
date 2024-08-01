package ac1

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 哈希表+两次遍历
// 用一个哈希表记录旧节点与新节点的映射关系
// 第一次遍历建立 Next 指针
// 第二次遍历建立 Random 指针
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := new(Node)
	old2New := map[*Node]*Node{head: newHead}
	for nNode, oNode := newHead, head; oNode != nil; nNode, oNode = nNode.Next, oNode.Next {
		nNode.Val = oNode.Val
		if oNode.Next != nil {
			next := new(Node)
			nNode.Next = next
			old2New[oNode.Next] = next
		}
	}
	for nNode, oNode := newHead, head; oNode != nil; nNode, oNode = nNode.Next, oNode.Next {
		nNode.Random = old2New[oNode.Random]
	}
	return newHead
}
