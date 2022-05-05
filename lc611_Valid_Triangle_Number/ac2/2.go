package ac2

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		k := i + 1
		for j := i + 1; j < len(nums); j++ {
			for k < len(nums) && nums[k] < nums[i]+nums[j] {
				k++
			}
			res += max(k-1-j, 0)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}