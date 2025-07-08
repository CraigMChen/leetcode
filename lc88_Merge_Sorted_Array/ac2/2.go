package ac2

// 正向双指针
// 先把 nums1 中的所有元素向右移动 n 个位置
// 然后正向使用双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}
	i, j, k := n, 0, 0
	for i < m+n && j < n {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}
