package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	nodeQue := []*TreeNode{root}
	sumQue := []int{0}
	for len(nodeQue) > 0 {
		node := nodeQue[0]
		nodeQue = nodeQue[1:]
		sum := sumQue[0]
		sumQue = sumQue[1:]
		if node.Left == nil && node.Right == nil {
			ans += node.Val + sum*10
			continue
		}
		if node.Left != nil {
			nodeQue = append(nodeQue, node.Left)
			sumQue = append(sumQue, node.Val+sum*10)
		}
		if node.Right != nil {
			nodeQue = append(nodeQue, node.Right)
			sumQue = append(sumQue, node.Val+sum*10)
		}
	}
	return ans
}
