package sorts

func merge(arr []int, m int) {
	tmp := make([]int, len(arr))
	i, j, k := 0, m, 0
	for i < m && j < len(arr) {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}
	for i < m {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j < len(arr) {
		tmp[k] = arr[j]
		k++
		j++
	}
	for i := range arr {
		arr[i] = tmp[i]
	}
}

// MergeSort 归并排序
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)，非原地排序，需要 O(logn) 大小的栈帧空间
// 稳定排序
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	m := len(arr) / 2
	MergeSort(arr[:m])
	MergeSort(arr[m:])
	merge(arr, m)
	return
}
