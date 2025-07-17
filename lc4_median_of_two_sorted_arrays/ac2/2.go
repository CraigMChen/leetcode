package ac2

// 双指针
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2 := 0, 0
	i := -1
	mid1, mid2 := 0, 0
	l := len(nums1) + len(nums2)
	mid := l / 2
	for i < mid {
		mid1 = mid2
		if p2 == len(nums2) || p1 < len(nums1) && nums1[p1] <= nums2[p2] {
			mid2 = nums1[p1]
			p1++
		} else {
			mid2 = nums2[p2]
			p2++
		}
		i++
	}
	if l%2 == 0 {
		return float64(mid1+mid2) / 2
	}
	return float64(mid2)
}
