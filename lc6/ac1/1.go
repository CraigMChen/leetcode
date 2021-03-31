package ac1

func reverse(x int) int {
	var ans int64 = 0
	for x != 0 {
		ans = ans*10 + int64(x)%10
		x /= 10
	}
	if ans < -1<<31 || ans > 1<<31-1 {
		ans = 0
	}
	return int(ans)
}
