package ac2

// 快慢指针
func isHappy(n int) bool {
	slow, fast := squareSum(n), squareSum(squareSum(n))
	for fast != 1 && slow != fast {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))
	}
	return fast == 1
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
