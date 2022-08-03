package ac2

import "math/rand"

// 随机数法
// 期望时间复杂度 O(n)
// 空间复杂度 O(1)
func majorityElement(nums []int) int {
	for {
		n := rand.Intn(len(nums))
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == nums[n] {
				count++
			}
		}
		if count > len(nums)/2 {
			return nums[n]
		}
	}
}
