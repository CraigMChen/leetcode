package ac1

import "math/rand"

type Solution struct {
	rects [][]int
	sums []int
}

// 与lc528思路类似，这里计算的是矩形中点的个数的前缀和
func Constructor(rects [][]int) Solution {
	sums := make([]int, len(rects) + 1)
	for i := 0; i < len(rects); i++ {
		sums[i + 1] = sums[i] +
			(rects[i][2] - rects[i][0] + 1) * (rects[i][3] - rects[i][1] + 1)
	}
	return Solution{
		rects: rects,
		sums:  sums,
	}
}


func (this *Solution) Pick() []int {
	n := rand.Intn(this.sums[len(this.sums) - 1]) + 1
	l, r := 1, len(this.sums)
	for l < r {
		m := (r - l) >> 1 + l
		if this.sums[m] < n {
			l = m + 1
		} else {
			r = m
		}
	}

	x := rand.Intn(this.rects[l - 1][2] - this.rects[l - 1][0] + 1) + this.rects[l - 1][0]
	y := rand.Intn(this.rects[l - 1][3] - this.rects[l - 1][1] + 1) + this.rects[l - 1][1]
	return []int{x, y}
}
