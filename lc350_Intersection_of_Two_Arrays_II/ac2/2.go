package ac2

// map
func intersect(nums1 []int, nums2 []int) []int {
	counter := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		counter[nums1[i]]++
	}
	var ans []int
	for i := 0; i < len(nums2); i++ {
		if n := counter[nums2[i]]; n > 0 {
			counter[nums2[i]]--
			ans = append(ans, nums2[i])
		}
	}
	return ans
}
