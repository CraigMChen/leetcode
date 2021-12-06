package ac2

// 动态规划
// 类似方法1，分成单区间子段和双区间子段两种情况
// 对于双区间子段[0,i],[j,len(nums)-1]的和可以表示为
// (nums[0]+...+nums[i]) + (nums[j]+...+nums[len(nums)-1])
// 也可以表示为
// sum(nums) - (nums[i+1]+...+nums[j-1])
// 即nums数组的总和减去子数组的最小和
// 这样就把双区间子段最大和的问题转换成单区间子段最小和的问题，方法类似
// 前提双区间不能为空，即转化后的单区间不能是整个数组
// 可以做两遍动态规划，第一遍求nums[0,len(nums)-1)的单区间最小和
// 第二遍求nums[1,len(nums))的单区间最小和
func maxSubarraySumCircular(nums []int) int {
	// 先求单区间最大和
	res := -900000001
	cur := 0
	for _, n := range nums {
		if cur > 0 {
			cur += n
		} else {
			cur = n
		}
		res = getMax(res, cur)
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	// 求nums[0,len(nums)-1)的单区间最小和
	min := 900000001
	cur = 0
	for i := 0; i < len(nums)-1; i++ {
		if cur < 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		min = getMin(min, cur)
	}

	// 求nums[1,len(nums))的单区间最小和
	cur = 0
	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		min = getMin(min, cur)
	}
	return getMax(sum-min, res)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
