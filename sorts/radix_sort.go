package sorts

func getDigit(num, exp int) int {
	return num / exp % 10
}

func countingSortDigit(arr []int, exp int) {
	counter := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		digit := getDigit(arr[i], exp)
		counter[digit]++
	}
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		digit := getDigit(arr[i], exp)
		res[counter[digit]-1] = arr[i]
		counter[digit]--
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = res[i]
	}
}

// RadixSort 基数排序
// 时间复杂度 O(nk)
// 空间复杂度 O(n)，非原地排序
// 稳定排序
func RadixSort(arr []int) {
	maxN := arr[0]
	for i := 0; i < len(arr); i++ {
		maxN = max(maxN, arr[i])
	}
	k := 0
	for maxN > 0 {
		maxN /= 10
		k++
	}

	exp := 1
	for i := 0; i < k; i++ {
		countingSortDigit(arr, exp)
		exp *= 10
	}
}
