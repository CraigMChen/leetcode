package ac1

// 位运算
// 每个非答案元素都出现了三次，答案只出现了一次
// 那么每个非答案元素的第 i 个二进制位之和一定是 3 的倍数
// 将所有元素的第 i 个二进制位相加后对 3 取模，结果一定等于答案的第 i 个二进制位
// 注意数组中存在负数，需要考虑符号位
// 所以实现时需要将数字转换成 int32
// 时间复杂度 O(nlogC)，其中 C 表示数据范围，本题数字为 32 位，所以 C 为 32
// 空间复杂度 O(1)
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		sum := int32(0)
		for j := 0; j < len(nums); j++ {
			sum += (int32(nums[j]) >> i) & 1
		}
		ans |= (sum % 3) << i
	}
	return int(ans)
}
