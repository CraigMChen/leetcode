package ac2

import "math/rand"

// 快速排序
// 快速排序时，随机选择一个下标 pivot，若这个 pivot 恰好是倒数第 k 个，则该下标上的数就是答案
func findKthLargest(nums []int, k int) int {
	var quickSelect func(l, r, index int) int
	partition := func(l, r int) int {
		pivot := rand.Intn(r-l) + l
		nums[pivot], nums[l] = nums[l], nums[pivot]
		left, right := l, r-1
		for left < right {
			for left < right && nums[right] > nums[l] {
				right--
			}
			for left < right && nums[left] <= nums[l] {
				left++
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
			}
		}
		nums[l], nums[left] = nums[left], nums[l]
		return left
	}
	quickSelect = func(l, r, index int) int {
		pivot := partition(l, r)
		if pivot == index {
			return nums[pivot]
		}
		if pivot < index {
			return quickSelect(pivot+1, r, index)
		}
		return quickSelect(l, pivot, index)
	}
	return quickSelect(0, len(nums), len(nums)-k)
}
