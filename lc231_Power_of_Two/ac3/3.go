package ac3

// 判断 n 是否为给定范围内最大 2 的幂的约数
// 给定范围为 -1<<31 <= n <= 1<<31 - 1，最大的 2 的幂为 1<<30
// 时间复杂度 O(1)
// 空间复杂度 O(1)
func isPowerOfTwo(n int) bool {
	return n > 0 && 1<<30%n == 0
}
