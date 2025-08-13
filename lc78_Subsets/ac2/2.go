package ac2

// 位运算
// 可以用一个数的二进制位来表示 nums 数组中的元素是否被选择
// 从 0 到 2^len(nums) 枚举这个数，即可找出所有子集
func subsets(nums []int) [][]int {
	var ans [][]int
	for i := 0; i < 1<<len(nums); i++ {
		var cur []int
		for j := 0; 1<<j <= i; j++ {
			if i&(1<<j) == 1 {
				cur = append(cur, nums[j])
			}
		}
		ans = append(ans, cur)
	}
	return ans
}
