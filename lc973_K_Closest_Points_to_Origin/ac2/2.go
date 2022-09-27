package ac2

import "math/rand"

// 快速选择
func kClosest(points [][]int, k int) [][]int {
	var (
		distance    [][2]int
		res         [][]int
		quickSelect func(l, r, index int) int
	)
	for i := 0; i < len(points); i++ {
		distance = append(distance, [2]int{points[i][0]*points[i][0] + points[i][1]*points[i][1], i})
	}
	partition := func(l, r int) int {
		pivot := rand.Intn(r-l) + l
		distance[l], distance[pivot] = distance[pivot], distance[l]
		left, right := l, r-1
		for left < right {
			for left < right && distance[right][0] < distance[l][0] {
				right--
			}
			for left < right && distance[left][0] >= distance[l][0] {
				left++
			}
			if left < right {
				distance[left], distance[right] = distance[right], distance[left]
			}
		}
		distance[l], distance[left] = distance[left], distance[l]
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
	quickSelect(0, len(distance), len(distance)-k)
	for i := len(distance) - k; i < len(distance); i++ {
		res = append(res, points[distance[i][1]])
	}
	return res
}
