package ac1

// MyQueue
// 用一个栈存放 push 进来的元素
// 用另一个栈存放要 pop 和 peek 的元素
// 当进行 pop 和 peek 操作时，若 popStack 不为空，则直接出栈；否则将 pushStack 中的元素依次出栈再 push 到 popStack 中
type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) toPop() {
	for len(q.pushStack) > 0 {
		n := q.pushStack[len(q.pushStack)-1]
		q.pushStack = q.pushStack[:len(q.pushStack)-1]
		q.popStack = append(q.popStack, n)
	}
}

func (q *MyQueue) Push(x int) {
	q.pushStack = append(q.pushStack, x)
}

func (q *MyQueue) Pop() int {
	if len(q.popStack) == 0 {
		q.toPop()
	}
	n := q.popStack[len(q.popStack)-1]
	q.popStack = q.popStack[:len(q.popStack)-1]
	return n
}

func (q *MyQueue) Peek() int {
	if len(q.popStack) == 0 {
		q.toPop()
	}
	return q.popStack[len(q.popStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.pushStack) == 0 && len(q.popStack) == 0
}
