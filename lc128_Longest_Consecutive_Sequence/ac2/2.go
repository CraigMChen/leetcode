package ac2

// 哈希表保存 nums 中数字的集合
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}

	ans := 0
	for num := range set {
		if _, ok := set[num-1]; ok { // 若存在比当前值更小的数，跳过
			continue
		}
		l := 1
		n := num
		for _, ok := set[n+1]; ok; _, ok = set[n+1] { // 不断寻找比当前值大 1 的数
			n++
			l++
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}
