package ac3

// 栈+堆
// 用小顶堆来维护最小值
// Push 操作的时间复杂度为 O(logn)
// Pop 操作的时间复杂度为 O(n)
// 远不如前两种解法

type MinStack struct {
	stack []int
	heap  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.heap = append(this.heap, val)
	for i := len(this.heap) - 1; i > 0; {
		parent := (i - 1) >> 1
		if this.heap[parent] <= this.heap[i] {
			break
		}
		this.heap[parent], this.heap[i] = this.heap[i], this.heap[parent]
		i = parent
	}
}

func (this *MinStack) Pop() {
	target := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	index := 0
	for ; this.heap[index] != target; index++ {
	}
	this.heap[index], this.heap[len(this.heap)-1] =
		this.heap[len(this.heap)-1], this.heap[index]
	this.heap = this.heap[:len(this.heap)-1]
	for i := index; i<<1+1 < len(this.heap); {
		l, r, m := i<<1+1, i<<1+2, i
		if l < len(this.heap) && this.heap[m] >= this.heap[l] {
			m = l
		}
		if r < len(this.heap) && this.heap[m] >= this.heap[r] {
			m = r
		}
		if m == i {
			break
		}
		this.heap[i], this.heap[m] = this.heap[m], this.heap[i]
		i = m
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.heap[0]
}
