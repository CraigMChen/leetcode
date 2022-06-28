package ac1

const MOD int = 1e9 + 7

// 二分查找
// 设 preSum 为数组 nums 的前缀和数组
// prePreSum 为数组 preSum 的前缀和数组
// 题目要求的是 nums 数组的所有子数组的和经排序后的数组第 left 个元素到第 right 个元素的和
// 设 nums 的所有子数组的和排序后的数组为 subSums，它的前 k 个元素的和为 sum(k)
// 则题意可转化为，求 sum(right) - sum(left-1)
// 对于 sum(k)，可以通过如下的方法求得
// 通过二分+双指针的方法求出数组 subSums 的第k个元素为 num
// sum(k) 可以分为 subSums 中所有严格小于 num 的元素与等于 num 的元素
// 先求所有严格小于 num 的元素的和
// 通过双指针的方法，求出所有满足和严格小于 num 的 nums 数组的区间 [i,j)
// 对每一个这样的区间 [i,j)，可以求出其子数组的和为 sum1 = prePreSum[j] - prePreSum[j] - preSum[i]) * (j-i)
// 将所有 j-i 累加起来即为和严格小于 num 的子数组数量 count
// 然后求等于 num 的元素的和
// sum2 = num * (k - count)
// 最后 sum = sum1 + sum2
func rangeSum(nums []int, n int, left int, right int) int {
	preSum := make([]int, len(nums)+1) // nums 的前缀和
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	prePreSum := make([]int, len(nums)+1) // preSum 的前缀和
	for i := 0; i < len(nums); i++ {
		prePreSum[i+1] = prePreSum[i] + preSum[i+1]
	}

	getCount := func(m int) int {
		// 双指针法求出 subSums 数组中 m 前面有多少个元素
		count := 0
		for i, j := 0, 1; i < len(preSum); i++ {
			// preSum[j] - preSum[i] <= m 表明有 j-i 个子数组的和小于等于 m
			for j < len(preSum) && preSum[j]-preSum[i] <= m {
				j++
			}
			j--
			count += j - i
		}
		return count
	}
	getSum := func(k int) int {
		// 二分查找求出 subSums 数组中第 k 个元素是多少
		// 找到第一个 m，使得 subNums 中小于等于 m 的元素至少有 k 个
		l, r := 0, preSum[len(preSum)-1]
		for l < r {
			m := (r-l)>>1 + l
			if getCount(m) >= k {
				r = m
			} else {
				l = m + 1
			}
		}
		num := l

		sum, count := 0, 0
		// 双指针法求出严格小于 num 的区间 [i,j)
		for i, j := 0, 1; i < len(preSum); i++ {
			for j < len(preSum) && preSum[j]-preSum[i] < num {
				j++
			}
			j--
			count += j - i
			// 表达式 prePreSum[j] - prePreSum[i] - preSum[i] * (j-i) 的结果为数组 nums 区间 [i, j) 的所有元素组成子数组的和组成的数组的元素和
			sum = (sum%MOD + (prePreSum[j]-prePreSum[i]-preSum[i]*(j-i))%MOD) % MOD
		}
		// 求出所有等于 num 的 subSums 的元素和
		sum = (sum%MOD + ((k-count)*num)%MOD) % MOD
		return sum
	}
	return getSum(right) - getSum(left-1)
}
