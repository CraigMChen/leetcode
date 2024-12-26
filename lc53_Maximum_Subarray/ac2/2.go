package ac2

// 分治
// 区间 [l, r) 中， m 为 中点
// 则左子区间为 [l, m)，右子区间为 [m, r)
// 设：
// lSum 为区间 [l, r) 中以下标 l 为左端点的最大子数组的和
// rSum 为区间 [l, r) 中以下标 r-1 为右端点的最大子数组的和
// mSum 为区间 [l, r) 中的最大子数组的和
// iSum 为区间 [l, r) 中所有元素的和
// 则：
// lSum = max( 左子区间的 lSum, 左子区间的 iSum + 右子区间的 lSum)
// rSum = max( 右子区间的 rSum, 右子区间的 iSum + 左子区间的 rSum)
// mSum = max( 左子区间的 mSum, 右子区间的 mSum, 左子区间的 rSum + 右子区间的 lSum)
// iSum = 左子区间的 iSum + 右子区间的 iSum
func maxSubArray(nums []int) int {
	_, _, ans, _ := getXSum(nums, 0, len(nums))
	return ans
}

func getXSum(nums []int, l, r int) (int, int, int, int) {
	if l == r-1 {
		return nums[l], nums[l], nums[l], nums[l]
	}
	m := (r-l)>>1 + l
	lSum1, rSum1, mSum1, iSum1 := getXSum(nums, l, m)
	lSum2, rSum2, mSum2, iSum2 := getXSum(nums, m, r)
	lSum := max(lSum1, iSum1+lSum2)
	rSum := max(rSum2, iSum2+rSum1)
	mSum := max(max(mSum1, mSum2), rSum1+lSum2)
	iSum := iSum1 + iSum2
	return lSum, rSum, mSum, iSum
}
