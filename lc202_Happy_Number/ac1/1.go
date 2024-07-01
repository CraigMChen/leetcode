package ac1

// 哈希表
func isHappy(n int) bool {
	existed := make(map[int]struct{})
	for _, ok := existed[n]; !ok; _, ok = existed[n] {
		existed[n] = struct{}{}
		n = squareSum(n)
		if n == 1 {
			return true
		}
	}
	return false
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
