package ac1

// 大顶堆
func findKthLargest(nums []int, k int) int {
	var maxHeapify func(int, int)
	maxHeapify = func(size, i int) {
		l, r, max := i<<1+1, i<<1+2, i
		if l < size && nums[l] > nums[max] {
			max = l
		}
		if r < size && nums[r] > nums[max] {
			max = r
		}
		if max != i {
			nums[i], nums[max] = nums[max], nums[i]
			maxHeapify(size, max)
		}
	}
	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(len(nums), i)
	}

	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(i, 0)
	}
	return nums[0]
}
