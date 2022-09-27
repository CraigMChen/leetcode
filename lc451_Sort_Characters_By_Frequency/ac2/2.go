package ac2

// 桶排序
// 统计每个字母出现次数
// 对字母出现次数进行桶排序
// 根据桶排序结果构造答案
func frequencySort(s string) string {
	var (
		counter  = make(map[byte]int)
		bucket   = make(map[int][]byte)
		maxCount int
		res      []byte
	)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
		count := counter[s[i]]
		if count > maxCount {
			maxCount = count
		}
	}

	for k, v := range counter {
		bucket[v] = append(bucket[v], k)
	}
	for i := maxCount; i > 0; i-- {
		v, ok := bucket[i]
		if !ok {
			continue
		}
		for j := 0; j < len(v); j++ {
			for k := 0; k < i; k++ {
				res = append(res, v[j])
			}
		}
	}
	return string(res)
}
