package ac1

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	y := 0
	for tmp != 0 {
		y = y * 10 + tmp % 10
		tmp /= 10
	}
	return x == y
}
