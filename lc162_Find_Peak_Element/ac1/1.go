package ac1

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		if r == l + 1 {
			return l
		}
		m := (r-l)>>1 + l
		if m != r-1 {
			if nums[m] < nums[m+1] {
				l = m + 1
			} else {
				r = m
			}
		} else if m != l {
			if nums[m] < nums[m-1] {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	return l
}
