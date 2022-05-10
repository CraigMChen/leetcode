package ac1

// 二分查找
// 从[1,max(nums)]中暴力查找，对于每个结果i，可以用O(n)的时间计算出需要多少次操作
func minimumSize(nums []int, maxOperations int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	l, r := 1, max+1
	for l < r {
		m := (r-l)>>1 + l
		if getOperations(nums, m) > maxOperations {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func getOperations(nums []int, max int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= max {
			continue
		}
		n := nums[i] / max
		if n*max == nums[i] {
			n--
		}
		res += n
	}
	return res
}
