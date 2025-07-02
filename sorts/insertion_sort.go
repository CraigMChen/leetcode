package sorts

// InsertionSort 插入排序
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)，原地排序
// 稳定排序
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		index := i
		for j := 0; j < i; j++ {
			if arr[j] > arr[index] {
				index = j
				break
			}
		}
		if index == i {
			continue
		}
		tmp := arr[i]
		for j := index; j <= i; j++ {
			arr[j], tmp = tmp, arr[j]
		}
	}
}
