package ac2

// 位运算
// 用 n&(-n) 求出用 n 的最低位 1 表示的数
// 若这个数等于 n，说明 n 的二进制表示中最低位是唯一一个 1
// 那么 n 一定是 2 的幂
// 时间复杂度 O(1)
// 空间复杂度 O(1)
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(-n) == n
}
