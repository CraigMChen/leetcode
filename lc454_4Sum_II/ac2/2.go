package ac2

import "sort"

// 排序，二分
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var sum []int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum = append(sum, nums1[i]+nums2[j])
		}
	}
	sort.Ints(sum)
	ans := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			s := -nums3[i] - nums4[j]
			lb := binSearch(sum, s, false)
			ub := binSearch(sum, s, true)
			if lb < len(sum) && sum[lb] == s {
				ans += ub - lb
			}
		}
	}
	return ans
}

func binSearch(nums []int, target int, upper bool) int {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] < target || (upper && nums[m] <= target) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
