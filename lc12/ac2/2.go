package ac2

import "strings"

var roman = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
var romanMap = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

func intToRoman(num int) string {
	ans := ""
	for i := len(roman) - 1; i >= 0; i-- {
		if num < romanMap[i] {
			continue
		}
		n := num / romanMap[i]
		ans += strings.Repeat(roman[i], n)
		num -= n * romanMap[i]
		if num == 0 {
			break
		}
	}
	return ans
}
