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

var nilValue = 1001

// 位运算 优化
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []byte
	convert := func(n int) {
		var neg byte = 0
		if n < 0 {
			neg = 1 << 7
			n *= -1
		}
		low := byte(n & (1<<8 - 1))
		top := byte(n>>8) | neg
		res = append(res, top, low)
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			convert(nilValue)
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
	var ser []int
	for i := 0; i < len(data); i += 2 {
		top, low := data[i], data[i+1]
		neg := 1
		if top>>7 == 1 {
			neg = -1
			top &= 1<<7 - 1
		}
		n := int(top)<<8 | int(low)
		n *= neg
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
			Val:   n,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}
