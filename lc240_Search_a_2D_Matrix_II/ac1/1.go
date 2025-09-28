package ac1

// 二分查找
// 对每一行进行二分查找
// 时间复杂度 O(mlogn)
// 空间复杂度 O(1)
func searchMatrix(matrix [][]int, target int) bool {
	binSearch := func(nums []int, target int) bool {
		l, r := 0, len(nums)
		for l < r {
			m := (r-l)>>1 + l
			if nums[m] == target {
				return true
			} else if nums[m] < target {
				l = m + 1
			} else {
				r = m
			}
		}
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if binSearch(matrix[i], target) {
			return true
		}
	}
	return false
}
