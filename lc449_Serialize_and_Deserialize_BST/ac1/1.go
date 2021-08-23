package ac1

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// 深搜
// 保存前序遍历结果（用"x"表示空节点，保存所有空节点的信息），用","分隔
// 数字直接转成字符串
func (this *Codec) serialize(root *TreeNode) string {
	var (
		arr []string
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			arr = append(arr, "x")
			return
		}
		arr = append(arr, strconv.FormatInt(int64(node.Val), 10))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(arr, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	i := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if arr[i] == "x" || i >= len(arr) {
			i++
			return nil
		}
		node := &TreeNode{}
		val, _ := strconv.ParseInt(arr[i], 10, 64)
		i++
		node.Val = int(val)
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}
