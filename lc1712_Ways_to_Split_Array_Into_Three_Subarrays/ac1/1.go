package ac1

const MOD int = 1e9 + 7

// 二分查找
// 分别从下标x,y处将数组分割为三段：[:x], [x:y], [y:]
// 枚举x，二分查找y
// 设数组前缀和和preSum
// 分割后第一段的和为 sum1 = preSum[x-1]
// 第二段的和为 sum2 = preSum[y-1]-preSum[x-1]
// 第三段的和为 sum3 = preSum[n-1]-preSum[y-1]
// 对于每个x，分别进行两次二分查找
// 第一次：找到第一个使 sum2 >= sum1 的 y，结果为y1
// 第二次：找到第一个使 sum3 < sum2 的 y，结果为y2
// 则从 y1 到 y2 - 1 的 y 都是可行的分割方案
func waysToSplit(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	res := 0
	for i := 1; i < len(nums)-1; i++ {
		l, r := i+1, len(nums)
		for l < r {
			m := (r-l)>>1 + l
			if nums[m-1]-nums[i-1] >= nums[i-1] {
				r = m
			} else {
				l = m + 1
			}
		}
		x := l

		l, r = i+1, len(nums)
		for l < r {
			m := (r-l)>>1 + l
			if nums[len(nums)-1]-nums[m-1] < nums[m-1]-nums[i-1] {
				r = m
			} else {
				l = m + 1
			}
		}
		y := l - 1
		if y >= x {
			res += y - x + 1
		}
	}
	return res % MOD
}
