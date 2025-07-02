package sorts

// SelectionSort 选择排序
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
// 非稳定排序
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		m := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		arr[i], arr[m] = arr[m], arr[i]
	}
}
