package ac1

// 贪心
// 记录当前片段的起始位置和终止位置 start, end
// 遍历到 i 时，若 s[i] 最后一次出现的位置 pos[s[i]-'a'] 大于 end，则用该值更新 end
// 当 i == end 时说明一个片段结束了，片段长度为 end - start + 1，更新 start 和 end 继续遍历
func partitionLabels(s string) []int {
	pos := make([]int, 26)
	for i := 0; i < len(s); i++ {
		pos[s[i]-'a'] = i
	}
	var res []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if pos[s[i]-'a'] > end {
			end = pos[s[i]-'a']
		}
		if i == end {
			res = append(res, end-start+1)
			start, end = i+1, i+1
		}
	}
	return res
}
