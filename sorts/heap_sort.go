package sorts

func heapify(arr []int, i int) {
	for i<<1+1 < len(arr) {
		l, r, m := i<<1+1, i<<1+2, i
		if l < len(arr) && arr[l] > arr[m] {
			m = l
		}
		if r < len(arr) && arr[r] > arr[m] {
			m = r
		}
		if m == i {
			break
		}
		arr[m], arr[i] = arr[i], arr[m]
		i = m
	}
}

// HeapSort 堆排序
// 时间复杂度 O(nlogn)
// 空间复杂度 O(1)，原地排序
// 非稳定排序：在交换堆顶元素和堆底元素时，相等元素的相对位置可能发生变化
func HeapSort(arr []int) {
	for i := (len(arr) - 2) >> 1; i >= 0; i-- {
		heapify(arr, i)
	}
	
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr[:i], 0)
	}
}
