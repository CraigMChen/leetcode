package ac1

import "sort"

// 排序+双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
			} else if nums[j]+nums[k] > -nums[i] {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
