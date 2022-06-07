package ac1

// 二分查找
// 从区间[1,maxSum-n+2)中用二分查找搜索下标为index的元素的值
// 当下标为index的元素的值为m时，为了使数组的和最小，数组元素应该从下标index往两边逐渐递减至1
// 此时可以用O(1)的时间复杂度求出数组的和
func maxValue(n int, index int, maxSum int) int {
	getSum := func(m int) int {
		sum := 0
		// 求数组下标范围在区间[0,index]的元素最小和
		if index >= m-1 {
			sum += m*(m+1)/2 + index - m + 1
		} else {
			sum += (2*m - index) * (index + 1) / 2
		}
		// 求数组下标范围在区间(index,n)的元素最小和
		if n-index >= m {
			sum += m*(m-1)/2 + n - index - m
		} else {
			sum += (2*m - n + index) * (n - index - 1) / 2
		}
		return sum
	}

	l, r := 1, maxSum-n+2
	for l < r {
		m := (r-l)>>1 + l
		if getSum(m) > maxSum {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
