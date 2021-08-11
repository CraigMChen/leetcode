package ac2

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 计数
// 方法1的优化，可以将槽位累加到一个计数器上
func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	sons := 1
	for i := 0; i < len(nodes); i++ {
		sons--
		if sons < 0 {
			return false
		}
		if nodes[i] != "#" {
			sons+=2
		}
	}
	return sons == 0
}
