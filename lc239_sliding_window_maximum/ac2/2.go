package ac2

// 单调队列
// 维护一个单调队列，其中的元素是数组 nums 的下标，队列中的下标对应的值在队列中是严格单调递减的
// 滑动窗口每向右移动一位时，如果当前元素大于等于队尾元素，则不断从队尾出队直到当前元素小于队尾元素
// 那么队首元素对应的值一定是最大值，但是不一定在滑动窗口中
// 不断从队首出队元素，直到队首元素在滑动窗口中
// 此时队首的元素对应的值就是滑动窗口的最大值
// 时间复杂度：O(n)
// 空间复杂度：O(k)
func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0, k+1)
	ans := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		if i >= k-1 {
			ans[i-k+1] = nums[queue[0]]
		}
	}
	return ans
}
