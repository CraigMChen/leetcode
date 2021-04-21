package ac1

// map
func intersection(nums1 []int, nums2 []int) []int {
	counter := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		counter[nums1[i]] = 1
	}

	var ans []int
	for i := 0; i < len(nums2); i++ {
		if isExist := counter[nums2[i]]; isExist == 1 {
			counter[nums2[i]]++
			ans = append(ans, nums2[i])
		}
	}
	return ans
}
