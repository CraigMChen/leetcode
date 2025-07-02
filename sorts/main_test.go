package sorts

import (
	"math/rand"
	"sort"
	"testing"
)

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCommon(t *testing.T) {
	var arr []int
	for i := 0; i < 100; i++ {
		//n := rand.Intn(1000) * int(math.Pow(-1, float64(rand.Intn(100)%2)))
		n := rand.Intn(1000)
		arr = append(arr, n)
	}
	
	arrSorted := make([]int, len(arr))
	copy(arrSorted, arr)
	sort.Ints(arrSorted)

	selectionSortArr := make([]int, len(arr))
	copy(selectionSortArr, arr)
	SelectionSort(selectionSortArr)
	if !isSliceEqual(arrSorted, selectionSortArr) {
		t.Error("SelectionSort is incorrect: ", selectionSortArr)
	}

	bubbleSortArr := make([]int, len(arr))
	copy(bubbleSortArr, arr)	
	BubbleSort(bubbleSortArr)
	if !isSliceEqual(arrSorted, bubbleSortArr) {
		t.Error("BubbleSort is incorrect: ", bubbleSortArr)
	}

	insertionSortArr := make([]int, len(arr))
	copy(insertionSortArr, arr)
	InsertionSort(insertionSortArr)
	if !isSliceEqual(arrSorted, insertionSortArr) {
		t.Error("InsertionSort is incorrect: ", insertionSortArr)
	}

	quickSortArr := make([]int, len(arr))
	copy(quickSortArr, arr)
	QuickSort(quickSortArr)
	if !isSliceEqual(arrSorted, quickSortArr) {
		t.Error("QuickSort is incorrect: ", quickSortArr)
	}
	
	mergeSortArr := make([]int, len(arr))
	copy(mergeSortArr, arr)
	MergeSort(mergeSortArr)
	if !isSliceEqual(arrSorted, mergeSortArr) {
		t.Error("MergeSort is incorrect: ", mergeSortArr)
	}
	
	heapSortArr := make([]int, len(arr))
	copy(heapSortArr, arr)
	HeapSort(heapSortArr)
	if !isSliceEqual(arrSorted, heapSortArr) {
		t.Error("HeapSort is incorrect: ", heapSortArr)
	}
	
	bucketSortArr := make([]int, len(arr))
	copy(bucketSortArr, arr)
	BucketSort(bucketSortArr, 10)
	if !isSliceEqual(arrSorted, bucketSortArr) {
		t.Error("BucketSort is incorrect: ", bucketSortArr)
	}
	
	countingSortArr := make([]int, len(arr))
	copy(countingSortArr, arr)
	CountingSort(countingSortArr)
	if !isSliceEqual(arrSorted, countingSortArr) {
		t.Error("CountingSort is incorrect: ", countingSortArr)
	}
	
	radixSortArr := make([]int, len(arr))
	copy(radixSortArr, arr)
	RadixSort(radixSortArr)
	if !isSliceEqual(arrSorted, radixSortArr) {
		t.Error("RadixSort is incorrect: ", radixSortArr)
	}
}
