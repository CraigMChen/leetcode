package ac1

// 动态规划
// 设 dp[i] 为第 i 个丑数，dp[1] = 1
// p2, p3, p5 初始时都为1
// dp[i] = min(dp[p2] * 2, dp[p3] * 3, dp[p5] * 5)
// 分别比较 dp[i] 与 dp[p2] * 2, dp[p3] * 3 和 dp[p5] * 5 的大小
// 若 dp[i] 与其相等，则相应的指针加1
//func nthUglyNumber(n int) int {
//	p2, p3, p5 := 1, 1, 1
//	dp := make([]int, n+1)
//	dp[1] = 1
//	for i := 2; i <= n; i++ {
//		dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
//		if dp[i] == dp[p2]*2 {
//			p2++
//		}
//		if dp[i] == dp[p3]*3 {
//			p3++
//		}
//		if dp[i] == dp[p5]*5 {
//			p5++
//		}
//	}
//	return dp[n]
//}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var uglyNumber []int

func init() {
	uglyNumber = make([]int, 1690)
	uglyNumber[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < 1690; i++ {
		r2 := uglyNumber[p2] * 2
		r3 := uglyNumber[p3] * 3
		r5 := uglyNumber[p5] * 5
		uglyNumber[i] = min(min(r2, r3), r5)
		if uglyNumber[i] == r2 {
			p2++
		}
		if uglyNumber[i] == r3 {
			p3++
		}
		if uglyNumber[i] == r5 {
			p5++
		}
	}
}

func nthUglyNumber(n int) int {
	return uglyNumber[n-1]
}