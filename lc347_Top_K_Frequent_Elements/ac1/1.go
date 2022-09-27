package ac1

// 堆排序
// 先统计每个元素出现的次数
// 出现次数数组中前 K 大的元素即为前 K 个高频元素
func topKFrequent(nums []int, k int) []int {
	var (
		counter    = make(map[int]int)
		heap       [][2]int
		maxHeapify func(int, int)
		res        []int
	)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}
	for k, v := range counter {
		heap = append(heap, [2]int{k, v})
	}

	maxHeapify = func(size int, n int) {
		l, r, max := n<<1+1, n<<1+2, n
		if l < size && heap[l][1] > heap[max][1] {
			max = l
		}
		if r < size && heap[r][1] > heap[max][1] {
			max = r
		}
		if max != n {
			heap[max], heap[n] = heap[n], heap[max]
			maxHeapify(size, max)
		}
	}
	for i := len(heap)/2 - 1; i >= 0; i-- {
		maxHeapify(len(heap), i)
	}

	size := len(heap)
	for i := 0; i < k; i++ {
		heap[0], heap[size-1] = heap[size-1], heap[0]
		size--
		maxHeapify(size, 0)
	}
	for i := len(heap) - k; i < len(heap); i++ {
		res = append(res, heap[i][0])
	}
	return res
}
