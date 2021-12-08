package ac1

// 动态规划
// 设pos[i]表示以下标i结尾的乘积为正数的最大长度
// neg[i]表示以下标i结尾的乘积为负数的最大长度，则
// 当nums[i] > 0 时
// pos[i] = pos[i-1]+1
// neg[i] = neg[i-1]+1 (neg[i-1] != 0)
//        = 0 (neg[i-1] == 0)
// 当nums[i] < 0 时
// pos[i] = neg[i-1]+1 (neg[i-1] != 0)
//        = 0 (neg[i-1] == 0)
// neg[i] = pos[i-1]+1
// 当nums[i] == 0 时
// pos[i] = neg[i] = 0
func getMaxLen(nums []int) int {
	res := 0
	pos, neg := 0, 0
	for _, n := range nums {
		if n > 0 {
			pos += 1
			if neg != 0 {
				neg++
			}
		} else if n == 0 {
			pos = 0
			neg = 0
		} else {
			tempPos, tempNeg := 0, 0
			if neg != 0 {
				tempPos = neg + 1
			}
			tempNeg = pos + 1
			pos, neg = tempPos, tempNeg
		}
		if pos > res {
			res = pos
		}
	}
	return res
}
