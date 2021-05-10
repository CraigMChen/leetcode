package ac1

import "sort"

// 二分
// 求出每个word的最小字母的出现频次并排序
// 对于每个query，通过二分搜索找到words中频次大于该频次的下标
func numSmallerByFrequency(queries []string, words []string) []int {
	fre := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		fre[i] = getMinCharacterFrequency(words[i])
	}
	sort.Ints(fre)

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		f := getMinCharacterFrequency(queries[i])
		l, r := 0, len(words)
		for l < r {
			m := (r - l) >> 1 + l
			if fre[m] <= f {
				l = m + 1
			} else {
				r = m
			}
		}
		ans[i] = len(words) - l
	}
	return ans
}

func getMinCharacterFrequency(str string) int {
	counter := make(map[int]int, 26)
	min := 27
	for i := 0; i < len(str); i++ {
		if int(str[i] - 'a') < min {
			min = int(str[i] - 'a')
		}
		counter[int(str[i] - 'a')]++
	}
	return counter[min]
}
