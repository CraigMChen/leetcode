package ac1

var roman2int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ans := 0
	l := len(s)
	pre := 0
	for i := l - 1; i >= 0; i-- {
		cur := roman2int[s[i]]
		if cur < pre {
			ans -= cur
		} else {
			ans += cur
		}
		pre = cur
	}
	return ans
}
