package ac2

import (
	"sort"
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

// 不用分隔符，用2个字节表示数字
func (this *Codec) serialize(root *TreeNode) string {
	const byt = 0xff
	var (
		res []byte
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, byte(node.Val&byt))
		res = append(res, byte(node.Val>>8))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return string(res)
}

func (this *Codec) deserialize(data string) *TreeNode {
	var (
		preorder []int
		build    func(pre, in []int) *TreeNode
	)
	for i := 0; i < len(data); i += 2 {
		a := int(data[i])
		b := int(data[i+1]) << 8
		preorder = append(preorder, a+b)
	}
	inorder := make([]int, len(preorder))
	copy(inorder, preorder)
	sort.Ints(inorder)

	build = func(pre, in []int) *TreeNode {
		if len(pre) == 0 {
			return nil
		}
		root := &TreeNode{Val: pre[0]}
		i := 0
		for ; i < len(in) && in[i] != root.Val; i++ {
		}
		root.Left = build(pre[1:i+1], in[:i])
		root.Right = build(pre[i+1:], in[i+1:])
		return root
	}
	return build(preorder, inorder)
}
