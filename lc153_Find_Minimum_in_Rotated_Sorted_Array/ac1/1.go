package ac1

// 二分
func findMin(nums []int) int {
	l, r := 0, len(nums)
	ans := 5001
	for l < r {
		m := (r - l) >> 1 + l
		if nums[l] == nums[m] {
			if nums[l] < ans {
				ans = nums[l]
			}
			l++
		} else if nums[m] == nums[r - 1] {
			if nums[m] < ans {
				ans = nums[m]
			}
			r--
		} else if nums[m] < nums[r - 1] {
			if nums[m] < ans {
				ans = nums[m]
			}
			r = m
		} else if nums[l] < nums[m] {
			if nums[l] < ans {
				ans = nums[l]
			}
			l = m + 1
		}
	}
	return ans
}
