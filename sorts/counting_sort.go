package sorts

// CountingSort 计数排序
// 时间复杂度 O(n+m)，m 为数组中最大的数
// 空间复杂度 O(n+m)，非原地排序
// 稳定排序
func CountingSort(arr []int) {
	arrMax := arr[0]
	for _, n := range arr {
		arrMax = max(arrMax, n)
	}
	
	counter := make([]int, arrMax+1)
	for _, n := range arr {
		counter[n]++
	}
	
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}
	
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res[counter[arr[i]]-1] = arr[i]
		counter[arr[i]]--
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = res[i]
	}
}
