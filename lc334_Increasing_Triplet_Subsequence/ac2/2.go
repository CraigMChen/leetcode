package ac2

const MAXINT = int(^uint(0) >> 1)

// 贪心
// 用 first 表示三元组第一个数，second 表示三元组第二个数
// 初始时， first = nums[0], second = MAXINT
// 遍历 nums，若 nums[i] < first 则更新 first
// 若 first < nums[i] < second 则更新 second
// 若 nums[i] > second 说明找到符合条件的三元组
func increasingTriplet(nums []int) bool {
	first := nums[0]
	second := MAXINT
	for i := 1; i < len(nums); i++ {
		if nums[i] < first {
			first = nums[i]
		} else if nums[i] > first && nums[i] < second {
			second = nums[i]
		} else if nums[i] > second {
			return true
		}
	}
	return false
}
