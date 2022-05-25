package ac1

import (
	"math/rand"
	"time"
)

type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{preSum: w}
}

// 二分查找
// 把权重想象成一条线段，所有线段首位相连
// 在连接后的整条线段上随机取一点
// 该点落在线段上的几率即为符合权重的几率
// 对权重求前缀和preSum
// 在[0,preSum[n-1])中随机取一个数
// 用二分查找找到第一个大于n的元素下标即为答案
func (this *Solution) PickIndex() int {
	n := rand.Intn(this.preSum[len(this.preSum)-1])
	l, r := 0, len(this.preSum)
	for l < r {
		m := (r-l)>>1 + l
		if this.preSum[m] > n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
