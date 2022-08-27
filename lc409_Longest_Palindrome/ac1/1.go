package ac1

// 贪心
// 出现次数为偶数的字符一定可以组成回文串，直接累加到结果中
// 出现次数为奇数的字符去掉一个，也可以组成回文串，将出现次数减一后的值累加到结果中
// 特例：出现次数为奇数的字符可以放在中间，从而最大化结果，此时将结果加 1 即可
func longestPalindrome(s string) int {
	counter := make(map[int32]int)
	for _, b := range s {
		counter[b]++
	}

	sum := 0
	odd := 0
	for _, v := range counter {
		if v&1 == 1 {
			odd++
			sum += v - 1
		} else {
			sum += v
		}
	}
	if odd > 0 {
		sum++
	}
	return sum
}
