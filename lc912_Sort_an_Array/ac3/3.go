package ac3

// 堆排序
// 堆的性质（大顶堆为例，小顶堆相反）：
// 每个节点的值大于它的所有子节点的值
// 堆顶元素一定为最大值
// 可以将数组元素建堆
// 从最后一个非叶子节点开始（最后一个节点下标为 len(nums)-1，最后一个非叶子节点是它的父节点，即 len(nums)/2 - 1）
// 不断对该节点之前的节点堆化
// 得到一个堆
// 此时堆顶为最大值
// 堆顶元素出堆，即把堆顶元素与最后一个元素交换，然后堆的元素数量减一
// 再对堆顶元素做堆化
// 得到的堆顶元素为第 2 大的值，再次将堆顶元素出堆，循环直到堆的元素为 0
// 最终得到的即为升序数组
func sortArray(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums[:i], 0)
	}
	return nums
}

// 自节点 i 开始向下堆化，使以节点 i 为根节点的子树成为一个堆
func maxHeapify(nums []int, i int) {
	for {
		// 不断检查当前节点与其左右儿子的大小，并与较大的节点做交换
		l, r, m := i<<1+1, i<<1+2, i
		if l < len(nums) && nums[l] > nums[m] {
			m = l
		}
		if r < len(nums) && nums[r] > nums[m] {
			m = r
		}
		if m == i { // 由于堆是自底向上建立的，如果当前节点已经是最大值，说明已经堆化结束
			break
		}
		nums[i], nums[m] = nums[m], nums[i]
		i = m
	}
}
