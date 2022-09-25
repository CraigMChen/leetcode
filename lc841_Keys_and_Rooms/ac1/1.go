package ac1

// 广搜
func canVisitAllRooms(rooms [][]int) bool {
	var (
		visited = make([]bool, len(rooms))
		count   = 0
		queue   = []int{0}
	)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if !visited[node] {
			visited[node] = true
			count++
			for i := 0; i < len(rooms[node]); i++ {
				if !visited[rooms[node][i]] {
					queue = append(queue, rooms[node][i])
				}
			}
		}
	}
	return count == len(rooms)
}
