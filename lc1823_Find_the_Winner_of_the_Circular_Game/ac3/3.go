package ac3

// 迭代版约瑟夫环
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func findTheWinner(n int, k int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (res + k) % i
	}
	return res + 1
}
