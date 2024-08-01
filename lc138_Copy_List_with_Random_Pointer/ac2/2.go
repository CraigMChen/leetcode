package ac2

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 哈希表+递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func copyRandomList(head *Node) *Node {
	old2New := make(map[*Node]*Node)

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := old2New[node]; ok {
			return n
		}
		newNode := new(Node)
		old2New[node] = newNode
		newNode.Val = node.Val
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}
	return deepCopy(head)
}
