package ac2

// 滑动窗口
// nums1不动，遍历nums2与nums1的对齐位置
// nums2不动，遍历nums1与nums2的对齐位置
// 找出每个对齐位置中的最大公共子数组，最大值即为答案
func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	l1 := len(nums1)
	l2 := len(nums2)
	for i := 0; i < l1; i++ {
		l := maxLength(nums1, nums2, 0, i)
		if l > ans {
			ans = l
		}
	}
	for i := 0; i < l2; i++ {
		l := maxLength(nums1, nums2, i, 0)
		if l > ans {
			ans = l
		}
	}
	return ans
}

func maxLength(nums1, nums2 []int, i, j int) int {
	k := 0
	ret := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] != nums2[j] {
			k = 0
		} else {
			k++
			if k > ret {
				ret = k
			}
		}
		i++
		j++
	}
	return ret
}
