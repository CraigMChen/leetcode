package ac1

// 如果 x 不能到 y+1，那么 [x, y] 之间任意一点都不能到 y+1
// 利用这个特性，可以从 0 开始，不断检查是否能走完
// 若不能走完，不需要检查下一个下标，而是检查第一个无法到达的下标
// 只需遍历一遍数组，复杂度为 O(n)
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		curGas := gas[i]
		j := 1
		for ; j <= len(gas); j++ {
			curGas -= cost[(i+j-1)%len(gas)]
			if curGas < 0 {
				break
			}
			curGas += gas[(i+j)%len(gas)]
		}
		if curGas >= 0 {
			return i
		}
		i += j
	}
	return -1
}
