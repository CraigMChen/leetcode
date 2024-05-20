package ac1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先用 dfs 找到两个目标点的位置和深度
// 不断将深度较大的点向上移动，直到两个点相同，一边移动一边填充字符串
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var (
		dfs     func(root *TreeNode, target, depth int) (*TreeNode, int)
		parents = make(map[*TreeNode]*TreeNode)
	)
	dfs = func(root *TreeNode, target, depth int) (*TreeNode, int) {
		if root == nil {
			return nil, depth
		}
		if root.Val == target {
			return root, depth
		}
		if root.Right != nil {
			parents[root.Right] = root
		}
		if root.Left != nil {
			parents[root.Left] = root
		}
		right, rd := dfs(root.Right, target, depth+1)
		if right != nil {
			return right, rd
		}
		return dfs(root.Left, target, depth+1)
	}

	start, sDepth := dfs(root, startValue, 0)
	dest, dDepth := dfs(root, destValue, 0)

	var (
		ans     []byte
		reverse []byte
	)
	for start != dest {
		if sDepth >= dDepth {
			ans = append(ans, 'U')
			start = parents[start]
			sDepth--
		} else {
			parent := parents[dest]
			if dest == parent.Right {
				reverse = append(reverse, 'R')
			} else {
				reverse = append(reverse, 'L')
			}
			dest = parent
			dDepth--
		}
	}

	for i := len(reverse) - 1; i >= 0; i-- {
		ans = append(ans, reverse[i])
	}
	return string(ans)
}
