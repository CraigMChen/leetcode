package ac1

// 归并
// 先将两个数组合并成一个，然后求中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))
	i, j, k := 0, 0, 0
	for i < len(nums1) || j < len(nums2) {
		if j == len(nums2) || i < len(nums1) && nums1[i] <= nums2[j] {
			merged[k] = nums1[i]
			i++
		} else {
			merged[k] = nums2[j]
			j++
		}
		k++
	}

	mid := len(merged) / 2
	if len(merged)%2 == 0 {
		return float64(merged[mid-1]+merged[mid]) / 2
	} else {
		return float64(merged[mid])
	}
}
