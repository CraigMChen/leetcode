package ac1

import "sort"

// 二分查找
// 双重二分
// 首先对数组进行排序
// 用二分查找在区间[0,max(arr)]中查找value
// 对于每一个value，用二分查找找到arr中第一个大于等于value的下标i，可以用前缀和求出替换后数组的和
// 最后需要比较 upper_bound 与 lower_bound（即 lower_bound-1）哪个更符合要求
func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	preSum := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		preSum[i+1] = preSum[i] + arr[i]
	}

	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}

	getSum := func(m int) int {
		l, r := 0, len(arr)
		for l < r {
			mid := (r-l)>>1 + l
			if arr[mid] > m {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return preSum[l] + m*(len(arr)-l)
	}

	l, r := 0, arr[len(arr)-1]+1
	for l < r {
		m := (r-l)>>1 + l
		if getSum(m) > target {
			r = m
		} else {
			l = m + 1
		}
	}

	if abs(getSum(l)-target) < abs(getSum(l-1)-target) {
		return l
	}
	return l - 1
}
