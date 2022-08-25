package ac1

// 前/后缀积
// 先预处理后缀积，遍历 nums 时动态维护前缀积
// 前缀积和后缀积之积即为结果
// 结果数组的空间可以用于维护后缀积，省空间
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums)+1)
	res[len(res)-1] = 1
	for i := len(res) - 2; i >= 0; i-- {
		res[i] = res[i+1] * nums[i]
	}
	preProduct := 1
	for i := 0; i < len(nums); i++ {
		res[i] = preProduct * res[i+1]
		preProduct *= nums[i]
	}
	return res[:len(res)-1]
}
