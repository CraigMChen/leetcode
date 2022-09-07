package ac2

// 约瑟夫环（递归）
// f(n,k) 表示有 n 个人，每次第 k 个人淘汰，最终的获胜者的下标
// 有递推公式：
// f(n,k) = (f(n-1,k)+k)%n
// 当 n == 1 时，f(n,k) = 0，即只剩一个人的时候，获胜者的下标肯定是 0
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func findTheWinner(n int, k int) int {
	var josephus func(n, k int) int
	josephus = func(n, k int) int {
		if n == 1 {
			return 0
		}
		return (josephus(n-1, k) + k) % n
	}
	return josephus(n, k) + 1
}
