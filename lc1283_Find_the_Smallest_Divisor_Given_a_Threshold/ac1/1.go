package ac1

// 二分
// 暴力搜索除数，对每个除数n，可以用O(n)的复杂度求出和
func smallestDivisor(nums []int, threshold int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	getSum := func(m int) int {
		res := 0
		for i := 0; i < len(nums); i++ {
			res += (nums[i]-1)/m + 1
		}
		return res
	}
	l, r := 1, max+1
	for l < r {
		m := (r-l)>>1 + l
		if getSum(m) <= threshold {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
