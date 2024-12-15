package ac2

import "math/rand"

// 快速排序
func sortArray2(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	pivot := partition(nums, 0, len(nums))
	sortArray2(nums[:pivot])
	sortArray2(nums[pivot+1:])
	return nums
}

func partition(nums []int, l, r int) int {
	pivot := getPivot1(l, r)
	nums[pivot], nums[0] = nums[0], nums[pivot] // 先将哨兵移动到最左侧
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j] >= nums[0] {
			j--
		}
		for i < j && nums[i] <= nums[0] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[0], nums[i] = nums[i], nums[0] // 将哨兵移动到分界处，此时哨兵左侧的数一定小于等于它，右侧的数一定大于等于它
	return i                            // 返回哨兵的下标
}

// 取随机下标作为哨兵
func getPivot1(l, r int) int {
	return rand.Intn(r-l) + l
}

// 取 nums[l], nums[r], nums[(r+1)/2] 的中间数下标作为哨兵
func getPivot2(nums []int, l, r int) int {
	m := (l + r) / 2
	if nums[l] >= nums[m] && nums[l] <= nums[r-1] ||
		nums[l] >= nums[r-1] && nums[l] <= nums[m] {
		return l
	}
	if nums[m] >= nums[l] && nums[m] <= nums[r-1] ||
		nums[m] >= nums[r-1] && nums[m] <= nums[l] {
		return m
	}
	return r - 1
}

func main() {
	sortArray2([]int{5, 2, 3, 1})
}
