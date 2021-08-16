package ac1

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

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
	res []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var (
		res []int
		dfs func([]*NestedInteger)
	)
	dfs = func(list []*NestedInteger) {
		for _, integer := range list {
			if integer.IsInteger() {
				res = append(res, integer.GetInteger())
			} else {
				dfs(integer.GetList())
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator{res}
}

func (this *NestedIterator) Next() int {
	ans := this.res[0]
	this.res = this.res[1:]
	return ans
}

func (this *NestedIterator) HasNext() bool {
	return len(this.res) != 0
}
