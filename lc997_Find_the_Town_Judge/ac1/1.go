package ac1

// 找到出度为 0 且入度为 n-1 的节点
func findJudge(n int, trust [][]int) int {
	in := make([]int, n+1)
	out := make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		out[trust[i][0]]++
		in[trust[i][1]]++
	}
	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}
