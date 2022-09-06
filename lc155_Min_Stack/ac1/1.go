package ac1

type MinStack struct {
	vals   []int
	preMin []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	if len(this.preMin) == 0 {
		this.preMin = append(this.preMin, val)
	} else {
		min := this.preMin[len(this.preMin)-1]
		if val < min {
			min = val
		}
		this.preMin = append(this.preMin, min)
	}
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}
	n := len(this.vals)
	this.vals = this.vals[:n-1]
	this.preMin = this.preMin[:n-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.preMin[len(this.preMin)-1]
}
