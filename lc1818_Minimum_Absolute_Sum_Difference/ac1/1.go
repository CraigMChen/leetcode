package ac1

import "sort"

const MOD int = 1e9 + 7

// 二分查找
// sum = |nums1[0]-nums2[0]| + ... + |nums1[n-1]-nums2[n-1]|
// 用nums1[i]替换nums1[j]，得到的结果为
// sum - |nums1[i]-nums2[i]| + |nums1[j]-nums2[i]|
// 要使结果最小，只要后面部分最大即可，即求
// max(|nums1[i]-nums2[i]| - |nums1[j]-nums2[i]|)
// 对于每个i，用二分查找找出nums1中与nums2[i]最接近的元素
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	absolute := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	sum := 0
	nums0 := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		sum += absolute(nums1[i], nums2[i])
		nums0[i] = nums1[i]
	}
	sort.Ints(nums0)

	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	max := 0
	for i := 0; i < len(nums1); i++ {
		l, r := 0, len(nums1)
		for l < r {
			m := (r-l)>>1 + l
			if nums0[m] >= nums2[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		if l < len(nums1) {
			max = getMax(max, absolute(nums1[i], nums2[i])-absolute(nums0[l], nums2[i]))
		}
		if l > 0 {
			max = getMax(max, absolute(nums1[i], nums2[i])-absolute(nums0[l-1], nums2[i]))
		}
	}
	return (sum - max) % MOD
}
