package sorts

// BucketSort 桶排序
// 时间复杂度 O(n+k), k 为桶的数量
// 空间复杂度 O(n+k)
// 是否稳定取决于桶内排序算法是否稳定
func BucketSort(arr []int, bucketCount int) {
	maxN, minN := arr[0], arr[0]
	for _, n := range arr {
		maxN = max(maxN, n)
		minN = min(minN, n)
	}

	buckets := make([][]int, bucketCount)
	// 向上取整
	bucketSize := (maxN - minN + bucketCount - 1) / bucketCount
	for _, n := range arr {
		index := (n - minN) / bucketSize
		buckets[index] = append(buckets[index], n)
	}

	for i := 0; i < bucketCount; i++ {
		InsertionSort(buckets[i])
	}

	i := 0
	for _, bucket := range buckets {
		for _, n := range bucket {
			arr[i] = n
			i++
		}
	}
}
