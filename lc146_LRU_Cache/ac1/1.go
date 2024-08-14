package ac1

type Node struct {
	parent   *Node
	next     *Node
	key, val int
}

type LRUCache struct {
	capacity   int
	head, tail *Node
	set        map[int]*Node
}

func (this *LRUCache) push(node *Node) {
	this.set[node.key] = node
	this.tail.parent.next = node
	node.parent = this.tail.parent
	node.next = this.tail
	this.tail.parent = node
}

func (this *LRUCache) delete(node *Node) {
	delete(this.set, node.key)
	node.parent.next = node.next
	node.next.parent = node.parent
}

func (this *LRUCache) refresh(node *Node) {
	this.delete(node)
	this.push(node)
}

func Constructor(capacity int) LRUCache {
	head := new(Node)
	tail := new(Node)
	head.next = tail
	tail.parent = head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		set:      make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.set[key]
	if !ok {
		return -1
	}
	this.refresh(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.set[key]
	if ok {
		node.val = value
		this.refresh(node)
		return
	}
	this.push(&Node{key: key, val: value})
	if len(this.set) > this.capacity {
		this.delete(this.head.next)
	}
}
