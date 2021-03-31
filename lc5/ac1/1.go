package ac1

// 暴力。每2n-2个分为一组
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l := len(s)
	ans := make([]byte, l)
	count := 0
	group := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; ; j++ {
			id1 := i + group*j
			id2 := id1 + 2*(numRows-1-i)
			if id1 >= l {
				break
			}
			ans[count] = s[id1]
			count++
			if i != 0 && i != numRows-1 && id2 > 0 && id2 < l { // 第一行或最后一行
				ans[count] = s[id2]
				count++
			}
		}
	}
	return string(ans)
}
