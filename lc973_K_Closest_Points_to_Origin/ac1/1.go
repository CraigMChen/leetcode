package ac1

// 堆排序
func kClosest(points [][]int, k int) [][]int {
	var (
		distance   [][2]int
		res        [][]int
		minHeapify func(int, int)
	)
	for i := 0; i < len(points); i++ {
		distance = append(distance, [2]int{points[i][0]*points[i][0] + points[i][1]*points[i][1], i})
	}
	minHeapify = func(size, n int) {
		l, r, min := n<<1+1, n<<1+2, n
		if l < size && distance[l][0] < distance[min][0] {
			min = l
		}
		if r < size && distance[r][0] < distance[min][0] {
			min = r
		}
		if min != n {
			distance[n], distance[min] = distance[min], distance[n]
			minHeapify(size, min)
		}
	}
	for i := len(distance)/2 - 1; i >= 0; i-- {
		minHeapify(len(distance), i)
	}

	size := len(distance)
	for i := 0; i < k; i++ {
		distance[0], distance[size-1] = distance[size-1], distance[0]
		size--
		minHeapify(size, 0)
	}
	for i := len(distance) - k; i < len(distance); i++ {
		res = append(res, points[distance[i][1]])
	}
	return res
}
