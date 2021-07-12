package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 找前驱节点
// 不需要递归，直接从根节点开始遍历，节省栈空间
func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			left := cur.Left
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			right := cur.Right
			pre.Right = right
			cur.Left, cur.Right = nil, left
		}
		cur = cur.Right
	}
}
