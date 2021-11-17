package ac1

// 动态规划
func minCostClimbingStairs(cost []int) int {
	x := cost[0]
	y := cost[1]
	z := 0
	for i := 2; i < len(cost); i++ {
		if x > y {
			z = y + cost[i]
		} else {
			z = x + cost[i]
		}
		x = y
		y = z
	}
	if x > y {
		return y
	}
	return x
}
