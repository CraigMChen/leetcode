package ac1

// 分别从左往右和从右往左遍历
// 序号比上一个大，就多发 1 颗糖，否则就只发 1 颗
// 得到的两个结果数组中分别去较大值相加
func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}

	res := max(1, candies[len(ratings)-1])
	curCandy := 1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			curCandy++
		} else {
			curCandy = 1
		}
		res += max(candies[i], curCandy)
	}
	return res
}
