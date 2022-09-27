package ac1

import "sort"

// 快速选择
// 记录每个字母出现的次数
// 对字母出现次数数组进行快速选择
// 根据排序后的数组构造答案
func frequencySort(s string) string {
	var (
		counter = make(map[byte]int)
		arr     [][2]int
		res     []byte
	)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	for k, v := range counter {
		arr = append(arr, [2]int{int(k), v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i][1]; j++ {
			res = append(res, byte(arr[i][0]))
		}
	}
	return string(res)
}
