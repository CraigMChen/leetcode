package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeQue := []*TreeNode{root}
	valueQue := []int{1}
	res := 0
	for len(nodeQue) != 0 {
		width := valueQue[len(valueQue)-1] - valueQue[0] + 1
		if width > res {
			res = width
		}
		var newNodeQue []*TreeNode
		var newValueQue []int
		for i := range nodeQue {
			node := nodeQue[i]
			value := valueQue[i]
			if node.Left != nil {
				newNodeQue = append(newNodeQue, node.Left)
				newValueQue = append(newValueQue, value*2)
			}
			if node.Right != nil {
				newNodeQue = append(newNodeQue, node.Right)
				newValueQue = append(newValueQue, value*2+1)
			}
		}
		nodeQue = newNodeQue
		valueQue = newValueQue
	}
	return res
}
