package ac1

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	tmp1, tmp2 := nums[0], nums[1]
	p := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != tmp1 || nums[i] != tmp2 {
			tmp1, tmp2 = tmp2, nums[i]
			nums[p] = nums[i]
			p++
		} else {
			tmp1, tmp2 = tmp2, nums[i-1]
		}
	}
	return p
}
