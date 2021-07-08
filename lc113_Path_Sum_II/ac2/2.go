package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广搜
// 用哈希表保存每个节点的父节点，找到目标叶子节点后就可以用哈希表还原整个路径
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	nodeQue := []*TreeNode{root}
	numQue := []int{0}
	parents := make(map[*TreeNode]*TreeNode)

	getPath := func(node *TreeNode) (path []int) {
		for node != nil {
			path = append(path, node.Val)
			node = parents[node]
		}
		for i, j := 0, len(path) - 1; i < j; {
			path[i], path[j] = path[j], path[i]
			i++
			j--
		}
		return path
	}

	for len(nodeQue) != 0 {
		node := nodeQue[0]
		nodeQue = nodeQue[1:]
		num := numQue[0]
		numQue = numQue[1:]
		if node.Left == nil && node.Right == nil && node.Val + num == targetSum {
			ans = append(ans, getPath(node))
		}
		if node.Left != nil {
			nodeQue = append(nodeQue, node.Left)
			numQue = append(numQue, num + node.Val)
			parents[node.Left] = node
		}
		if node.Right != nil {
			nodeQue = append(nodeQue, node.Right)
			numQue = append(numQue, num + node.Val)
			parents[node.Right] = node
		}
	}
	return ans
}
