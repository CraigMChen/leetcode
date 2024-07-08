package ac1

import "fmt"

func summaryRanges(nums []int) []string {
	var ans []string
	for i, j := 0, 0; j < len(nums); {
		for ; j < len(nums) && nums[j]-nums[i] == j-i; j++ {
		}
		if i == j-1 {
			ans = append(ans, fmt.Sprintf("%d", nums[i]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
		}
		i = j
	}
	return ans
}
