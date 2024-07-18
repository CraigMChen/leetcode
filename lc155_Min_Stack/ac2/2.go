package ac2

// 单个栈
// 用一个变量 minVal 维护当前最小值
// 栈中维护入栈的值与当前最小值的差值
// 则出栈值 vTop = top + minVal (top > 0)
//                  = minVal (top <= 0)
// 出栈时，若当前栈顶值 top < 0，则需更新最小值 minVal = minVal + top

type MinStack struct {
	stack  []int
	minVal int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 { // 当栈为空时，最小值即为入栈值，差值为 0
		this.stack = append(this.stack, 0)
		this.minVal = val
		return
	}
	this.stack = append(this.stack, val-this.minVal)
	if val < this.minVal {
		this.minVal = val
	}
}

func (this *MinStack) Pop() {
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if val <= 0 { // 若栈顶值小于 0，说明此时出栈的值为最小值，需要更新 minVal
		this.minVal = this.minVal - val
	}
}

func (this *MinStack) Top() int {
	val := this.stack[len(this.stack)-1]
	if val > 0 { // 若栈顶值大于 0，说明上一个最小值与当前最小值是相同的
		val = val + this.minVal // 则可以推算出出栈值
	} else { // 若栈顶值小于 0，则出栈值即为当前最小值
		val = this.minVal
	}
	return val
}

func (this *MinStack) GetMin() int {
	return this.minVal
}
