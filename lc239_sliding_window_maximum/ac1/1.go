package ac1

// 优先队列（大顶堆）
// 将滑动窗口中的每个数和它的下标作为一个元素加入到大顶堆中
// 滑动窗口每次向右移动一位时，将该数入堆并自下而上堆化
// 堆顶的元素是最大值，但不一定在滑动窗口中
// 所以不断将堆顶元素出堆，直到堆顶元素在滑动窗口中
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)
func maxSlidingWindow(nums []int, k int) []int {
	heap := make([][2]int, 0, len(nums))
	ans := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		heap = append(heap, [2]int{nums[i], i})
		siftUp(heap, len(heap)-1)
		top := heap[0]
		for top[1] <= i-k {
			heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
			heap = heap[:len(heap)-1]
			siftDown(heap, 0)
			top = heap[0]
		}
		if i >= k-1 {
			ans[i-k+1] = top[0]
		}
	}
	return ans
}

func siftDown(nums [][2]int, i int) {
	for i < len(nums) {
		l, r, m := i<<1+1, i<<1+2, i
		if l < len(nums) && nums[l][0] > nums[m][0] {
			m = l
		}
		if r < len(nums) && nums[r][0] > nums[m][0] {
			m = r
		}
		if m == i {
			break
		}
		nums[m], nums[i] = nums[i], nums[m]
		i = m
	}
}

func siftUp(nums [][2]int, i int) {
	for i > 0 {
		p := (i - 1) >> 1
		if nums[p][0] > nums[i][0] {
			break
		}
		nums[i], nums[p] = nums[p], nums[i]
		i = p
	}
}
