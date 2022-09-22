package ac3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 记录子节点数量
// 用 map 记录以每个节点为根节点的子树中节点的数量
// 深搜，遍历到 node 时
// 若以 node.Left 为根节点的子树的节点数量等于 k-1，则 node.Val 即为所求
// 若以 node.Left 为根节点的子树的节点数量小于 k-1, 则目标值一定在右子树中，递归查找 node.Right，并将 k 更新为 k-counter[node.Left]-1
// 否则目标值一定在左子树中，递归查找 node.Left
func kthSmallest(root *TreeNode, k int) int {
	var (
		counter      = make(map[*TreeNode]int)
		getNodeCount func(*TreeNode) int
		findKth      func(*TreeNode) int
	)
	getNodeCount = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		count := getNodeCount(node.Left) + getNodeCount(node.Right) + 1
		counter[node] = count
		return count
	}
	findKth = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		if counter[node.Left] == k-1 {
			return node.Val
		}
		if counter[node.Left] < k-1 {
			k -= counter[node.Left] + 1
			return findKth(node.Right)
		}
		return findKth(node.Left)
	}
	getNodeCount(root)
	return findKth(root)
}
