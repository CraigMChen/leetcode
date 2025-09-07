package ac1

// 双哈希表
// 一个哈希表的键为访问次数，值为一个双向链表，链表节点保存 key, value 和使用次数
// 另一个哈希表的键为 key，值为双向链表的节点地址
// 需要用一个 minCount 变量来维护当前的最小使用次数
// 时间复杂度 O(1)
// 空间复杂度 O(n)

type node struct {
	key, val     int
	parent, next *node
	count        int
}

type cache struct {
	head, tail *node
}

func (c *cache) remove(n *node) {
	parent := n.parent
	next := n.next
	parent.next = next
	next.parent = parent
}

func (c *cache) add(n *node) {
	first := c.head.next
	c.head.next = n
	n.parent = c.head
	n.next = first
	first.parent = n
}

func newCache() *cache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.parent = head
	return &cache{
		head: head,
		tail: tail,
	}
}

type LFUCache struct {
	capacity int
	minCount int
	counts   map[int]*cache
	keys     map[int]*node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		counts:   make(map[int]*cache, capacity),
		keys:     make(map[int]*node, capacity),
	}
}

func (c *LFUCache) remove(n *node) {
	cache := c.counts[n.count]
	cache.remove(n)
	if cache.head.next == cache.tail {
		delete(c.counts, n.count)
		if n.count == c.minCount {
			c.minCount++
		}
	}
}

func (c *LFUCache) add(n *node) {
	cache := c.counts[n.count]
	if cache == nil {
		cache = newCache()
		c.counts[n.count] = cache
	}
	cache.add(n)
}

func (c *LFUCache) Get(key int) int {
	n, ok := c.keys[key]
	if !ok {
		return -1
	}

	c.remove(n)
	n.count++
	c.add(n)
	return n.val
}

func (c *LFUCache) Put(key int, value int) {
	n, ok := c.keys[key]
	if ok {
		c.remove(n)
		n.count++
		n.val = value
		c.add(n)
		return
	}

	if len(c.keys) >= c.capacity {
		last := c.counts[c.minCount].tail.parent
		c.remove(last)
		delete(c.keys, last.key)
	}

	n = &node{
		key:   key,
		val:   value,
		count: 1,
	}
	c.add(n)
	c.minCount = 1
	c.keys[key] = n
}
