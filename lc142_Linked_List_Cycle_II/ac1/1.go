package ac1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 用 map 记录已访问的节点，第一次遍历到的已访问节点即为答案
func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{})
	for p := head; p != nil; p = p.Next {
		if _, ok := visited[p]; ok {
			return p
		} else {
			visited[p] = struct{}{}
		}
	}
	return nil
}
