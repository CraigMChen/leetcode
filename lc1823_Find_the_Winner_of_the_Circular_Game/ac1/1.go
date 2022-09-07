package ac1

// 队列+模拟
// 用队列模拟游戏过程
// 将 1 到 n 分别入队
// 从队首取 k-1 个元素分别加入队尾，然后将队首元素取出，表示被淘汰的人
// 不断重复上面的步骤直到队列中只有一个人为止
// 时间复杂度 O(nk)，空间复杂度 O(n)
func findTheWinner(n int, k int) int {
	var queue []int
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}
	for len(queue) > 1 {
		for i := 0; i < k-1; i++ {
			queue = append(queue[1:], queue[0])
		}
		queue = queue[1:]
	}
	return queue[0]
}
