package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
// dfs求出所有子树的和并记录
// 遍历子树和的map，找出出现次数最多的子树和
func findFrequentTreeSum(root *TreeNode) []int {
	var (
		res []int
		dfs func(node *TreeNode) int
	)
	counts := make(map[int]int)
	max := -1
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			counts[node.Val]++
			if count := counts[node.Val]; count > max {
				max = count
			}
			return node.Val
		}
		sum := dfs(node.Left) + dfs(node.Right) + node.Val
		counts[sum]++
		if count := counts[sum]; count > max {
			max = count
		}
		return sum
	}
	dfs(root)
	for sum, count := range counts {
		if count == max {
			res = append(res, sum)
		}
	}
	return res
}
