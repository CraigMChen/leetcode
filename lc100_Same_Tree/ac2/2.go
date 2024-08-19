package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pQue := []*TreeNode{p}
	qQue := []*TreeNode{q}
	for len(pQue) > 0 && len(qQue) > 0 {
		pNode := pQue[len(pQue)-1]
		pQue = pQue[:len(pQue)-1]
		qNode := qQue[len(qQue)-1]
		qQue = qQue[:len(qQue)-1]
		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil || qNode == nil ||
			pNode.Val != qNode.Val {
			return false
		}
		pQue = append(pQue, pNode.Left, pNode.Right)
		qQue = append(qQue, qNode.Left, qNode.Right)
	}
	return len(pQue) == len(qQue)
}
