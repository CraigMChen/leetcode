package ac2

// 深搜
func canVisitAllRooms(rooms [][]int) bool {
	var (
		visited = make([]bool, len(rooms))
		count   = 0
		dfs     func(int)
	)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		count++
		for i := 0; i < len(rooms[node]); i++ {
			dfs(rooms[node][i])
		}
	}
	dfs(0)
	return count == len(rooms)
}
