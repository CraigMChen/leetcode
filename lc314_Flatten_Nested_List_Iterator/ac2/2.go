package ac2

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool {
	return true
}

func (n NestedInteger) GetInteger() int {
	return 0
}

func (n NestedInteger) SetInteger(int) {
}

func (n NestedInteger) Add(NestedInteger) {
}

func (n NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

type NestedIterator struct {
	stack [][]*NestedInteger
}

// Constructor
// 迭代
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := [][]*NestedInteger{nestedList}
	return &NestedIterator{stack}
}

func (this *NestedIterator) Next() int {
	que := this.stack[len(this.stack)-1]
	res := que[0].GetInteger()
	this.stack[len(this.stack)-1] = que[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) != 0 {
		que := this.stack[len(this.stack)-1]
		if len(que) == 0 {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		node := que[0]
		if node.IsInteger() {
			return true
		}
		this.stack[len(this.stack)-1] = que[1:]
		this.stack = append(this.stack, node.GetList())
	}
	return false
}
