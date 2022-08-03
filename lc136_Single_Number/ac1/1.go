package ac1

// 位运算
// 将数组中的数全部按位异或，结果就是答案
func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
