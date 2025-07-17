package ac3

// 二分查找
// 设两数组长度为 m 和 n
// 当 m + n 为奇数时，中位数的即为两数组合并后的第 (m+n)/2 个数
// 当 m + n 为偶数时，中位数的即为两数组合并后的第 (m+n)/2 和第 (m+n)/2+1 个数
// 所以这个问题可以转换为寻找两数组合并后的第 k 小的数
// 考虑两数组下标为 i = k/2-1 的数： x = nums1[i], y = nums2[i]
// 当 x <= y 时，小于 nums1[] 的数最多为 num1 数组中下标 x 之前的数和 nums2 数组中下标 y 之前的数
// 即 x <= y 时，小于 x 的数最多有 k/2-1 + k/2-1 = k-2 个
// 所以 nums1[i] 不可能是第 k 个数（因为比第 k 个数小的数有 k-1 个）
// 此时就可以排除 nums1 数组中下标为 i 以及之前的数
// 同理，当 x > y 时，可以排除 nums2 数组中下标为 i 之前的数
// 这里就可以使用二分查找算法，问题的规模每次缩小 k/2-1
// 边界条件：
// 当某个数组长度为0 时，返回另一个数组中第 k 个元素
// 当 k = 1 时，直接返回两数组第一个元素中较小的那个
// 当 k/2-1 越界时，选取数组中最后一个数字
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	k := l / 2
	if l%2 == 0 {
		return float64(findKth(nums1, nums2, k)+findKth(nums1, nums2, k+1)) / 2
	}
	return float64(findKth(nums1, nums2, k+1))
}

func findKth(nums1 []int, nums2 []int, k int) int {
	for {
		if len(nums1) == 0 {
			return nums2[k-1]
		}
		if len(nums2) == 0 {
			return nums1[k-1]
		}
		if k == 1 {
			return min(nums1[0], nums2[0])
		}
		n := k/2 - 1
		x, y := len(nums1)-1, len(nums2)-1
		if n < len(nums1) {
			x = n
		}
		if n < len(nums2) {
			y = n
		}
		if nums1[x] <= nums2[y] {
			nums1 = nums1[x+1:]
			k -= x + 1
		} else {
			nums2 = nums2[y+1:]
			k -= y + 1
		}
	}
}
