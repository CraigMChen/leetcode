package ac2

import "math/rand"

// 快速排序
func sortArray(nums []int) []int {
	partition := func(l, r int) int {
		pivot := rand.Intn(r-l) + l // 在 [l,r) 中随机选择一个下标
		nums[pivot], nums[l] = nums[l], nums[pivot]
		left, right := l, r-1
		// 将小于等于 nums[pivot] 的元素都放到 pivot 左边，大于它的放到右边
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
	var quickSort func(l, r int)
	quickSort = func(l, r int) {
		if l < r {
			pivot := partition(l, r)
			quickSort(l, pivot)
			quickSort(pivot+1, r)
		}
	}
	quickSort(0, len(nums))
	return nums
}
