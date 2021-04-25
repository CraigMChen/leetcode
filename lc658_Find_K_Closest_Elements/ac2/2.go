package ac2

import "sort"

// 按照与x的距离对arr进行排序，排序后的k个即为答案
func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i] - x) == abs(arr[j] - x) {
			return arr[i] < arr[j]
		}
		return abs(arr[i] - x) < abs(arr[j] - x)
	})
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
