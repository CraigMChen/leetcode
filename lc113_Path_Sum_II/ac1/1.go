package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜
func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		ans [][]int
		dfs func(node *TreeNode, num int, path []int)
	)

	dfs = func(node *TreeNode, num int, path []int) {
		if node == nil {
			return
		}
		left := num - node.Val
		newPath := append(path, node.Val)
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int{}, newPath...)) // 这里不能直接把newPath放入ans，因为切片复制时底层是同一个相关数组，后面遍历时可能会把ans中的结果改掉
			return
		}
		dfs(node.Left, left, newPath)
		dfs(node.Right, left, newPath)
	}
	dfs(root, targetSum, nil)
	return ans
}

// 深搜 优化
func pathSum2(root *TreeNode, targetSum int) [][]int {
	var (
		ans [][]int
		dfs func(node *TreeNode, num int)
		path []int
	)

	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		left := num - node.Val
		path = append(path, node.Val)
		defer func() { // TODO: 为什么用defer就是正确的，不用defer，放到后面就是错的？？？
			path = path[:len(path) - 1]
		}()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
		path = path[:len(path) - 1]
	}
	//dfs(root, targetSum)
	return ans
}
