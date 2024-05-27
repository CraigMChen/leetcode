package main

// 模拟
func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	curLen := 0
	curNum := 0
	l, r := 0, -1
	for i := 0; i < len(words); {
		if curLen+len(words[i])+curNum <= maxWidth {
			curLen += len(words[i])
			curNum++
			r = i
			i++
			continue
		}

		var line []byte
		space := maxWidth - curLen
		left := 0
		if curNum > 1 {
			space = (maxWidth - curLen) / (curNum - 1)
			left = (maxWidth - curLen) % (curNum - 1)
		}
		for j := l; j <= r; j++ {
			line = append(line, words[j]...)
			s := space
			if j == r && curNum > 1 {
				continue
			}
			if j-l < left {
				s++
			}
			for k := 0; k < s; k++ {
				line = append(line, ' ')
			}
		}
		ans = append(ans, string(line))
		curLen = 0
		curNum = 0
		l, r = i, i
	}
	var line []byte
	for i := l; i <= r; i++ {
		line = append(line, words[i]...)
		if i != len(words)-1 {
			line = append(line, ' ')
		}
	}
	n := len(line)
	for i := 0; i < maxWidth-n; i++ {
		line = append(line, ' ')
	}
	ans = append(ans, string(line))
	return ans
}

func main() {
	fullJustify([]string{
		"This",
		"is",
		"an",
		"example",
		"of",
		"text",
		"justification.",
	}, 16)
}
