package ac2

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

var nilValue int16 = 1001

// 位运算 优化
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []byte
	convert := func(n int) {
		low := byte(int16(n) & (1<<8 - 1))
		top := byte(int16(n) >> 8)
		res = append(res, top, low)
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			convert(int(nilValue))
			return
		}
		convert(node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return string(res)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var ser []int16
	for i := 0; i < len(data); i += 2 {
		top, low := data[i], data[i+1]
		n := int16(top)<<8 | int16(low)
		ser = append(ser, n)
	}
	var build func() *TreeNode
	i := 0
	build = func() *TreeNode {
		if i >= len(ser) {
			return nil
		}
		n := ser[i]
		i++
		if n == nilValue {
			return nil
		}
		return &TreeNode{
			Val:   int(n),
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}
