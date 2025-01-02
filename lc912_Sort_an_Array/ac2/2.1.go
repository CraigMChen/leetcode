package ac2

import "math/rand"

func sortArray2(nums []int) []int {
	quickSort(nums, 0, len(nums))
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r { // 数组长度小于 1 时停止
		return
	}
	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot)
	quickSort(nums, pivot+1, r)
}

func partition(nums []int, l, r int) int {
	pivot := getPivot1(nums, l, r)
	nums[l], nums[pivot] = nums[pivot], nums[l] // 先将哨兵换到最左侧
	i, j := l, r-1
	for i < j {
		// 由于哨兵放在最左侧，所以这里要先从右边开始动，这样最终边界点一定小于等于哨兵元素
		for i < j && nums[j] >= nums[l] { // 从右边开始找到第一个小于哨兵元素的值
			j--
		}
		for i < j && nums[i] <= nums[l] { // 从左边开始找到第一个大于哨兵元素的值
			i++
		}
		nums[i], nums[j] = nums[j], nums[i] // 将较大者放到右边，较小者放到左边
	}
	nums[i], nums[l] = nums[l], nums[i] // 将哨兵元素交换回分界点
	return i
}

func getPivot1(nums []int, l, r int) int { // 在 [l, r) 中随机选择一个元素作为哨兵
	return rand.Intn(r-l) + l
}

func getPivot2(nums []int, l, r int) int { // 在 nums[l], nums[m], nums[r] 中选择一个中间值作为哨兵
	m := (r-l)>>1 + l
	a, b, c := nums[l], nums[m], nums[r-1]
	if a > b && a < c || a > c && a < b {
		return l
	}
	if b > a && b < c || b > c && b < a {
		return m
	}
	return r - 1
}
