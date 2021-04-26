package ac1

// 动态规划
// dp[i][j]表示nums1[i:]与nums2[j:]的最长公共前缀的长度，状态转移方程为
// dp[i][j] = dp[i + 1][j + 1] + 1, nums1[i] == nums2[j]
// dp[i][j] = 0, nums1[i] != nums2[j]
func findLength(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	dp := make([][]int, l1)
	ans := 0
	for i := l1 - 1; i >= 0; i-- {
		dp[i] = make([]int, l2)
		for j := l2 - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				if i == l1-1 || j == l2-1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j+1] + 1
				}
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
