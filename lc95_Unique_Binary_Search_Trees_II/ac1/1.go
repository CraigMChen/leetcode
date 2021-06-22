package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归/回溯
// 构造一个递归函数，func build(start, end int) []*TreeNode
// 表示用从start到end的数构造所有可能的二叉搜索树
func generateTrees(n int) []*TreeNode {
	var build func(start, end int) []*TreeNode
	build = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil} // 注意这里不能返回nil
		}
		if start == end {
			return []*TreeNode{{Val: start}}
		}
		var allTrees []*TreeNode
		for i := start; i <= end; i++ {
			leftTrees := build(start, i - 1) // 递归构造左右子树
			rightTrees := build(i + 1, end)
			for j := 0; j < len(leftTrees); j++ { // 遍历所有左右子树集合
				for k := 0; k < len(rightTrees); k++ {
					root := &TreeNode{Val: i}
					root.Left = leftTrees[j]
					root.Right = rightTrees[k]
					allTrees = append(allTrees, root)
				}
			}
		}
		return allTrees
	}
	return build(1, n)
}
