package ac1

func removeElement(nums []int, val int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p++
		}
	}
	return p + 1
}
