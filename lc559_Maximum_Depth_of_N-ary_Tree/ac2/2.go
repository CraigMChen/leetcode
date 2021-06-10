package ac2

type Node struct {
	Val      int
	Children []*Node
}

// 广搜
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	que := []*Node{root}
	levelQue := []int{1}
	ans := 0
	for len(que) != 0 {
		node := que[0]
		que = que[1:]
		level := levelQue[0]
		levelQue = levelQue[1:]
		if level > ans {
			ans = level
		}
		for _, child := range node.Children {
			if child != nil {
				que = append(que, child)
				levelQue = append(levelQue, level + 1)
			}
		}
	}
	return ans
}
