package ac2

// 卡特兰数
// Catalan(n) = C(2n, n) - C(2n, n+2)
// 直接算会溢出，需要化简
// Catalan(n+1) = (4n+2)/(n+2)*Catalan(n)
func numTrees(n int) int {
	ans := 1
	for i := 1; i < n; i++ {
		ans = ans * (4 * i + 2) / (i + 2)
	}
	return ans
}