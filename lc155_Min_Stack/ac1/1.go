package ac1

// 辅助栈
// 用一个辅助栈来保存每次 push 后的最小值
// pop 时辅助栈也同步出栈

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) > 0 && val > this.minStack[len(this.minStack)-1] {
		val = this.minStack[len(this.minStack)-1]
	}
	this.minStack = append(this.minStack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
