package ac1

// 哈希表
// 长度为 n 的数组，其中没有出现的最小正整数一定在 [1,n+1] 范围内
// 用一个哈希表记录 nums 中大于 0 的数是否出现过
// 然后从 1 到 n 遍历哈希表，若其中有数字没有出现过，则该数字就是答案，否则 n+1 就是答案
// 为了满足常数级的空间复杂度，可以利用原数组
// 由于只需要考虑数组中 [1,n] 的数，可以将小于 0 的数都改成 n+1，这样数组中所有数都是大于 0 的
// 遍历数组中的数，对于 nums[i]，可以利用原数组都大于 0 的性质，将下标为 nums[i]-1 的元素变成负数
// 表示 nums[i] 这个数在数组中出现过，即 nums[nums[i]-1] = -nums[nums[i]-1]
// 此时 nums 数组中，小于 0 的元素 nums[i] 表示数 i+1 在数组中出现过，大于 0 表示没有出现过
// 最后再遍历一次数组，第一个出现的大于 0 的元素的下标 +1 即为答案，若所有元素都小于 0，则 n+1 为答案
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		x := nums[i]
		if x < 0 {
			x = -x
		}
		if x <= n && nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		}
	}
	ans := n + 1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			ans = i + 1
			break
		}
	}
	return ans
}
