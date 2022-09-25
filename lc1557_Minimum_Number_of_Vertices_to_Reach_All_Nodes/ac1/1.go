package ac1

// 寻找入度为 0 的节点
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	in := make([]int, n)
	for i := 0; i < len(edges); i++ {
		in[edges[i][1]]++
	}
	var res []int
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
