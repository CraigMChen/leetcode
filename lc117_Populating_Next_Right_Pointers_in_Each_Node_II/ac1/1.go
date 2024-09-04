package ac1

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 使用已建立的next数组，与lc116类似
func connect(root *Node) *Node {
	first := root // 当前层的第一个节点
	for first != nil {
		var next, pre *Node // next为下一层的第一个节点，pre为上一个节点
		for first != nil {
			for _, node := range []*Node{first.Left, first.Right} { // 遍历当前节点的左右儿子
				if node == nil {
					continue
				}
				if next == nil { // 找到下一层的第一个节点
					next = node
				}
				// 将上一个节点与当前节点连接
				if pre == nil {
					pre = node
				} else {
					pre.Next = node
					pre = pre.Next
				}
			}
			first = first.Next // 遍历当前层的下一个节点
		}
		first = next // 遍历下一层
	}
	return root
}

func connect2(root *Node) *Node {
	var next, prior *Node
	for node := root; node != nil; {
		for _, n := range []*Node{node.Left, node.Right} {
			if n == nil {
				continue
			}
			if prior != nil {
				prior.Next = n
			}
			if next == nil {
				next = n
			}
			prior = n
		}
		if node.Next != nil {
			node = node.Next
		} else {
			node = next
			next = nil
			prior = nil
		}
	}
	return root
}
