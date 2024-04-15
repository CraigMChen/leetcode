package ac2

// 翻转数组
func rotate(nums []int, k int) {
	reverse := func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	reverse(0, len(nums)-1)
	reverse(0, k%len(nums)-1)
	reverse(k%len(nums), len(nums)-1)
}
