package ac2

type Node struct {
	Val      int
	Children []*Node
}

// 迭代
// 假设根节点u有v1,v2,v3三个子节点
// 则该树的后续遍历为
// v1, [children of v1], v2, [children of v2], v3, [children of v3], u
// 将该序列反转得
// u, [children of v3]', v3, [children of v2]', v2, [children of v1]', v1
// 其中[]'表示序列的倒序
// 该结果和前序遍历类似，只是子节点的入栈顺序相反
// 将结果反转即为最终结果
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ans []int
	stack := []*Node{root}
	for len(stack) != 0 {
		l := len(stack)
		node := stack[l - 1]
		stack = stack[:l - 1]
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			stack = append(stack, child)
		}
	}
	return reverse(ans)
}

func reverse(nums []int) []int {
	var res []int
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, nums[i])
	}
	return res
}
