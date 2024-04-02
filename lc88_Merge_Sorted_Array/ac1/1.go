package ac1

// 双指针
// 分别从 nums1 的 m-1 下标和 nums2 的末尾开始遍历，每次将最大值写到 nums1 的末尾
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, k := m-1, n-1, m+n-1; k >= 0; k-- {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if j < 0 {
			nums1[k] = nums1[i]
			i--
		} else if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
