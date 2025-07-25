package ac1

// 动态规划
// 对于字符串A、B，有六种操作方式，实际上有部分是等价的：
// 在A中插入一个字符 = 在B中删除一个字符
// 在A中删除一个字符 = 在B中插入一个字符
// 在A中替换一个字符 = 在B中替换一个字符
// 所以实际上只有三种操作
// 设 dp[i][j] 表示字符串A的前i个字符和字符串B的前j个字符的编辑距离
// 对于第一种操作，有
// dp[i][j] = dp[i-1][j] + 1
// 对于第二种操作，有
// dp[i][j] = dp[i][j-1] + 1
// 对于第三种操作，有
// dp[i][j] = dp[i-1][j-1] + 1, (A[i-1] != B[j-1])
// dp[i][j] = dp[i-1][j-1], (A[i-1] == B[j-1])
// 即
// dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1] + 1), (A[i-1] != B[j-1])
// dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1]), (A[i-1] == B[j-1])
// 边界条件：
// 当 i == 0 时，dp[i][j] = j
// 当 j == 0 时，dp[i][j] = i
// 可以使用滚动数组优化空间
func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
	for i := 0; i <= len(word1); i++ {
		x := 0 // 表示 dp[i-1][j-1]
		for j := 0; j <= len(word2); j++ {
			tmp := dp[j]
			if i == 0 {
				dp[j] = j
			} else if j == 0 {
				dp[j] = i
			} else if word1[i-1] == word2[j-1] {
				dp[j] = min(dp[j]+1, dp[j-1]+1, x)
			} else {
				dp[j] = min(dp[j]+1, dp[j-1]+1, x+1)
			}
			x = tmp
		}
	}
	return dp[len(word2)]
}
