package ac3

import "sort"

// 排序
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
