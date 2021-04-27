package ac1

import "math/rand"

type Solution struct {
	preSum []int
}


func Constructor(w []int) Solution {
	preSum := make([]int, len(w))
	preSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		preSum[i] = preSum[i - 1] + w[i]
	}
	return Solution{preSum: preSum}
}


// 求出数组w的前缀和，则w中的每个元素都可以看作是数轴上的线段，在线段上随机取一个点，该点落在某个线段上的概率与线段的长度成正比
func (this *Solution) PickIndex() int {
	n := rand.Intn(this.preSum[len(this.preSum) - 1]) + 1
	l, r := 0, len(this.preSum)
	for l < r {
		m := (r - l) >> 1 + l
		if this.preSum[m] < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
