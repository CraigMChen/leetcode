package ac2

import "math/rand"

// 快速选择
func topKFrequent(nums []int, k int) []int {
	var (
		counter     = make(map[int]int)
		arr         [][2]int
		res         []int
		quickSelect func(l, r, index int) int
	)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}
	for k, v := range counter {
		arr = append(arr, [2]int{k, v})
	}

	partition := func(l, r int) int {
		pivot := rand.Intn(r-l) + l
		arr[l], arr[pivot] = arr[pivot], arr[l]
		left, right := l, r-1
		for left < right {
			for left < right && arr[right][1] > arr[l][1] {
				right--
			}
			for left < right && arr[left][1] <= arr[l][1] {
				left++
			}
			if left < right {
				arr[left], arr[right] = arr[right], arr[left]
			}
		}
		arr[l], arr[left] = arr[left], arr[l]
		return left
	}
	quickSelect = func(l, r, index int) int {
		pivot := partition(l, r)
		if pivot == index {
			return pivot
		}
		if pivot < index {
			return quickSelect(pivot+1, r, index)
		}
		return quickSelect(l, pivot, index)
	}
	quickSelect(0, len(arr), len(arr)-k)
	for i := len(arr) - k; i < len(arr); i++ {
		res = append(res, arr[i][0])
	}
	return res
}
