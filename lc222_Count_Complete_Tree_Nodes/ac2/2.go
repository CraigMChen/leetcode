package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二分，位运算
// 根节点为1，向左儿子移动则在下一位加个0，即 n << 1
// 向右儿子移动则在下一个加个1，即 n << 1 + 1
// 利用这个性质，可以在O(logn)的时间里判断是否存在n个节点
// 利用二分查找，从[1<<h, 1<<(h+1))范围中找出最小的不存在的节点数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h := 0 // 树的深度
	for node := root; node.Left != nil; node = node.Left {
		h++
	}
	if h == 0 { // 如果树的深度为0，说明只有根节点
		return 1
	}

	hasNode := func(k int) bool {
		node := root
		bits := 1 << (h - 1) // 从做到右，一次把每一位（除最高位外）取出
		for node != nil && bits > 0 {
			if k&bits == 0 { // 如果是0则向左走
				node = node.Left
			} else { // 否则向右走
				node = node.Right
			}
			// 下面这种写法是不对的
			//if k & bits == 1 {
			//	node = node.Right
			//} else {
			//	node = node.Left
			//}
			bits >>= 1
		}
		return node != nil
	}

	l, r := 1<<h, 1<<(h+1) // 树的节点个数范围为[1<<h, 1<<(h+1))
	for l < r {
		m := (r-l)>>1 + l
		if hasNode(m) { // 相当于二分查找最小的不存在的节点数
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
