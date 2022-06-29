package ac1

// 二分查找
// 先找到左边非递减区间[:left]
// 和右边非递减区间[right:]
// 然后遍历左边的非递减区间，对于左边的每个下标i，在右边找到第一个大于等于arr[i]的元素下标j，则要删除的区间[i+1,j)长度为 j-i-1
// 遍历右边的非递减区间，对于右边的每个下标i，在左边找到第一个大于arr[i]的元素的下标j，则要删除的区间[j,i)长度为 i-j
func findLengthOfShortestSubarray(arr []int) int {
	left := 1
	for left < len(arr) && arr[left] >= arr[left-1] {
		left++
	}
	right := len(arr) - 1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}
	if left > right {
		return 0
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var res int = 1e5 + 1
	for i := 0; i < left; i++ {
		l, r := right, len(arr)
		for l < r {
			m := (r-l)>>1 + l
			if arr[m] >= arr[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		res = min(res, l-i-1)
	}

	for i := right; i < len(arr); i++ {
		l, r := 0, left
		for l < r {
			m := (r-l)>>1 + l
			if arr[m] > arr[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		res = min(res, i-l)
	}
	return res
}
