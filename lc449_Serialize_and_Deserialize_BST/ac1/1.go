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
func (this *Codec) serialize(root *TreeNode) string {
	var (
		arr []string
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			arr = append(arr, "null")
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
		if arr[i] == "null" || i >= len(arr) {
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
