package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	cur   *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

// 迭代中序遍历
func (this *BSTIterator) Next() int {
	cur := this.cur
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}
	cur = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.cur = cur.Right
	return cur.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) != 0
}
