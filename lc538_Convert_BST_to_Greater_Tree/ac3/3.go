package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			suc := node.Right
			for suc.Left != nil && suc.Left != node {
				suc = suc.Left
			}
			if suc.Left == nil {
				suc.Left = node
				node = node.Right
			} else {
				suc.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}
