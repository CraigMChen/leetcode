package ac1

type MyLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	p := this.Head
	for i := 0; i < index && p != nil; i++ {
		p = p.Next
	}
	if p == nil {
		return -1
	}
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: this.Head,
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val: val,
	}
	if this.Head == nil {
		this.Head, this.Tail = node, node
		return
	}
	this.Tail.Next = node
	this.Tail = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	var (
		i      int
		p      = this.Head
		parent *Node
	)
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	for i = 0; i < index && p != nil; i++ {
		parent = p
		p = p.Next
	}
	if i == index && p == nil {
		this.AddAtTail(val)
		return
	}
	if p == nil {
		return
	}
	parent.Next = &Node{
		Val:  val,
		Next: p,
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	var (
		i      int
		p      = this.Head
		parent *Node
	)
	if index < 0 || this.Head == nil {
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
		return
	}
	for i = 0; i < index && p != nil; i++ {
		parent = p
		p = p.Next
	}
	if p == nil {
		return
	}
	if i == index && p.Next == nil {
		this.Tail = parent
		if parent != nil {
			this.Tail.Next = nil
		}
		return
	}
	parent.Next = p.Next
}
