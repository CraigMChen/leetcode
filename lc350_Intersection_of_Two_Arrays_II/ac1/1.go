package ac1

import "sort"

// 双指针
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var ans []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			ans = append(ans, nums1[i])
			i++
			j++
		}
	}
	return ans
}
