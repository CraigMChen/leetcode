package ac1

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	for node := root; node.Left != nil; node = node.Left {
		for par := node; par != nil; par = par.Next {
			par.Left.Next = par.Right
			if par.Next != nil {
				par.Right.Next = par.Next.Left
			}
		}
	}
	return root
}
