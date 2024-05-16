package ac1

// 原地解法
// 先翻转整个字符串，然后翻转每个单词，把单词按顺序放到切片头部，去掉多余空格
func reverseWords(s string) string {
	reverse := func(str []byte) {
		for i := 0; 2*i < len(str); i++ {
			str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
		}

	}

	b := []byte(s) // go 字符串不可修改，用 byte 模拟原地解法
	reverse(b)
	idx := 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			continue
		}
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(b[i:j])
		if idx != 0 {
			b[idx] = ' '
			idx++
		}
		for k := i; k < j; k++ {
			b[idx] = b[k]
			idx++
		}
		i = j
	}
	return string(b[:idx])
}
