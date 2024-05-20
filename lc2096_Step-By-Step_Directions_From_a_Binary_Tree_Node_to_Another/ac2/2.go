package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用 dfs 分别求出从根节点到 start 节点和 dest 节点的路径 sPath 和 dPath
// 去掉 sPath 和 dPath 的公共前缀
// 将 sPath 剩余部分全部改为 'U' 并与 dPath 剩余部分拼接
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var (
		dfs func(node *TreeNode, val int, path []byte) []byte
	)
	dfs = func(node *TreeNode, val int, path []byte) []byte {
		if node == nil {
			return nil
		}

		if node.Val == val {
			return path
		}
		p := dfs(node.Right, val, append(path, 'R'))
		if p != nil {
			return p
		}
		return dfs(node.Left, val, append(path, 'L'))
	}

	sPath := dfs(root, startValue, nil)
	dPath := dfs(root, destValue, nil)

	i := 0
	for ; i < len(sPath) && i < len(dPath) && sPath[i] == dPath[i]; i++ {
	}

	var ans []byte
	for j := i; j < len(sPath); j++ {
		ans = append(ans, 'U')
	}
	for j := i; j < len(dPath); j++ {
		ans = append(ans, dPath[j])
	}
	return string(ans)
}
