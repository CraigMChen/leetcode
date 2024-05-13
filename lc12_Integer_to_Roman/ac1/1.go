package ac1

import (
	"fmt"
)

var roman = [][]byte{
	{'M'},
	{'C', 'D'},
	{'X', 'L'},
	{'I', 'V'},
}

func intToRoman(num int) string {
	strs := make([]string, 4)
	for i := 3; i >= 0 && num != 0; i-- {
		n := num % 10
		if n == 4 {
			strs[i] =  string([]byte{roman[i][0], roman[i][1]})
		} else if n == 9 {
			strs[i] = string([]byte{roman[i][0], roman[i-1][0]})
		} else if n < 4 {
			strs[i] = repeat(roman[i][0], n)
		} else {
			strs[i] = fmt.Sprintf("%c%s", roman[i][1], repeat(roman[i][0], n-5))
		}
		num /= 10
	}
	return fmt.Sprintf("%s%s%s%s", strs[0], strs[1], strs[2], strs[3])
}

func repeat(b byte, count int) string {
	res := make([]byte, count)
	for i := 0; i < count; i++ {
		res[i] = b
	}
	return string(res)
}
