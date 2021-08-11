package ac1

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 栈
func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	sons := []int{1} // sons表示当前上一个合法节点还剩余的儿子节点槽位
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		if node == "," {
			continue
		}
		if len(sons) == 0 { // 如果没有槽位，且还有节点没有遍历完，则一定非法
			return false
		}
		top := sons[len(sons)-1]
		top--
		sons[len(sons)-1] = top // 不论遇到"#"还是数字,槽位都要减1
		if top == 0 {
			sons = sons[:len(sons)-1] // 当上一个合法节点的槽位变成0时，该合法节点就被消耗完了，从栈中移除
		}
		if node != "#" {
			sons = append(sons, 2) // 如果遇到的是数字，则提供一个新的合法节点，即2个槽位
		}
	}
	return len(sons) == 0
}
