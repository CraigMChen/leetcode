package ac1

// 二分
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (r - l) >> 1 + l
		if arr[m] < arr[m + 1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
