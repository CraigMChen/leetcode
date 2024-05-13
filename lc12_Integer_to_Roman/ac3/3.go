package ac3

var int2Roman = map[int]byte{
	1:    'I',
	5:    'V',
	10:   'X',
	50:   'L',
	100:  'C',
	500:  'D',
	1000: 'M',
}

func intToRoman(num int) string {
	var ans []byte
	for cur := 1000; cur > 0; cur /= 10 {
		if num < cur {
			continue
		}
		n := num / cur
		if n == 4 {
			ans = append(ans, int2Roman[cur], int2Roman[cur*5])
		} else if n == 9 {
			ans = append(ans, int2Roman[cur], int2Roman[cur*10])
		} else {
			if n >= 5 {
				ans = append(ans, int2Roman[cur*5])
				n -= 5
			}
			for i := 0; i < n; i++ {
				ans = append(ans, int2Roman[cur])
			}
		}
		num %= cur
	}
	return string(ans)
}
