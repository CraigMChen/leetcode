package sorts

import "math/rand"

func partition(arr []int, l, r int) int {
	pivot := rand.Intn(r-l) + l
	arr[l], arr[pivot] = arr[pivot], arr[l]
	i, j := l, r-1
	for i < j {
		for i < j && arr[j] >= arr[l] {
			j--
		}
		for i < j && arr[i] <= arr[l] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[l], arr[i] = arr[i], arr[l]
	return i
}

// QuickSort 快速排序
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)，原地排序，需要 O(n) 的栈帧空间
// 非稳定排序
func QuickSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	pivot := partition(arr, 0, len(arr))
	QuickSort(arr[:pivot])
	QuickSort(arr[pivot+1:])
}
