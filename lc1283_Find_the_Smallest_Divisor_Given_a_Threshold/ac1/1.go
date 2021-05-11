package ac1

import "math"

// 二分
// 在[1, 1e6+1]内二分搜索除数，对每个除数n，遍历数组nums求出和与threshold比较
func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, int(1e6+1)
	for l < r {
		m := (r - l) >> 1 + l
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += int(math.Ceil(float64(nums[i]) / float64(m)))
		}
		if sum > threshold {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
