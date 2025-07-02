package sorts

// BubbleSort 冒泡排序
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
// 稳定排序
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
