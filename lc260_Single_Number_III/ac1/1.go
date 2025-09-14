package ac1

// 位运算
// 计算 nums 中所有数字的异或和 xor
// 计算 xor & (-xor) 求得 xor 的最低位 1 所在的位 l
// 设答案中的两个数为 x1 和 x2
// xor = nums[0] ^ nums[1] ^ ... nums[len(nums)-1] = x1 ^ x2
// 那么 x1 和 x2 中，一定会有一个数的第 l 位为 1，另一个数的第 l 位为 0
// 将 nums 中的数分为两类：
// 一类为第 l 位为 1 的数
// 一类为第 l 位为 0 的数
// 那么相同的数一定被归为同一类
// x1 和 x2 一定被归为不同类
// 分别求出两类数的异或和，结果即为 x1 和 x2
func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	l := xor & (-xor)
	x1, x2 := 0, 0
	for _, num := range nums {
		if num&l == 0 {
			x1 ^= num
		} else {
			x2 ^= num
		}
	}
	return []int{x1, x2}
}
