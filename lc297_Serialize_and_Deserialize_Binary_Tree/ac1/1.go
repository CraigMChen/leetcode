package ac1

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

// 位运算
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var (
		nilNodeVal = 1001
		res        []byte
		dfs        func(*TreeNode)
	)
	// 把一个 int 拆成两个 byte，用最高位表示符号，1 表示负数，0 表示正数
	convert := func(n int) []byte {
		neg := 0
		if n < 0 {
			neg = 1 << 7
			n = -n
		}
		return []byte{byte(n>>8 | neg), byte(n & (1<<9 - 1))}
	}
	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, convert(nilNodeVal)...)
			return
		}
		res = append(res, convert(node.Val)...)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return string(res)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var arr []int
	for i := 0; i < len(data); i += 2 {
		neg := 1
		if data[i]&(1<<7) > 0 {
			neg = -1
		}
		arr = append(arr, (int(data[i]&(1<<7-1))<<8|int(data[i+1]))*neg)
	}
	var build func() *TreeNode
	n := 0
	build = func() *TreeNode {
		if n >= len(arr) || arr[n] > 1000 {
			n++
			return nil
		}
		val := arr[n]
		n++
		return &TreeNode{
			Val:   val,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}
