package ac1

// 二分查找
// 先找到左边非递减区间[:left+1]
// 和右边非递减区间[right:]
// 然后遍历左边的非递减区间，对于左边的每个下标i，在右边找到第一个大于等于arr[i]的元素下标j，则要删除的区间长度为 j-1-(i+1)+1
// 遍历右边你的非递减区间，对于右边的每个下标i，在左边找到第一个大于arr[i]的元素的下标j，则要删除的区间长度为 i-1-j+1
func findLengthOfShortestSubarray(arr []int) int {
	left := 0
	for left < len(arr)-1 && arr[left] <= arr[left+1] {
		left++
	}
	right := len(arr) - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if left >= right {
		return 0
	}

	var res int = 1e5 + 1
	for i := 0; i <= left; i++ {
		l, r := right, len(arr)
		for l < r {
			m := (r-l)>>1 + l
			if arr[m] >= arr[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		if l-1-i < res {
			res = l - 1 - i
		}
	}
	for i := right; i < len(arr); i++ {
		l, r := 0, left+1
		for l < r {
			m := (r-l)>>1 + l
			if arr[m] > arr[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		if i-1-l < res {
			res = i - l
		}
	}
	return res
}
