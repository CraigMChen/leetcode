package ac1

// 二分
// 用二分找到arr中x的lower_bound，用双指针向两边扩展
func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)
	for l < r {
		m := (r - l) >> 1 + l
		if arr[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}

	index := l
	if l == len(arr) || l != 0 && x - arr[l - 1] <= arr[l] - x {
		index = l - 1
	}
	i, j := index, index + 1

	for j - i < k {
		if i <= 0 {
			j++
		} else if j >= len(arr) {
			i--
		} else if x - arr[i - 1] <= arr[j] - x {
			i--
		} else {
			j++
		}
	}
	return arr[i:j]
}
