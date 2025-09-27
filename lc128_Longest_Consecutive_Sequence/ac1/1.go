package ac1

// 哈希表保存以每个数为最大值可以组成的最长连续子序列的长度
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func longestConsecutive(nums []int) int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]] = 1 // 初始状态，所有数可以由本身组成长度为 1 的序列
	}

	ans := 0
	for num, l := range counter {
		n := num - 1
		for counter[n] > 0 { // 不断寻找比当前数小 1 的数
			l += counter[n] // 将长度累加
			counter[n] = 0  // 查找过的数清空，避免重复查找
			n--
		}
		counter[num] = l // 更新当前长度
		if l > ans {
			ans = l
		}
	}
	return ans
}
