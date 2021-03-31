package ac2

import (
	"strconv"
)

func reverse(x int) int {
	isPositive := 1
	if x < 0 {
		isPositive = -1
		x = -x
	}
	s := strconv.FormatInt(int64(x), 10)
	l := len(s)
	ansString := make([]byte, l)
	for i := 0; i < l; i++ {
		ansString[i] = s[l-i-1]
	}
	ans, _ := strconv.ParseInt(string(ansString), 10, 64)
	if ans < -1<<31 || ans > 1<<31-1 {
		ans = 0
	}
	return int(ans) * isPositive
}
