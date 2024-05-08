package ac2

// 从左到右遍历
// 前一个同学分配的糖果数量为 pre
// 记录当前同学所在递增或递减序列的长度 inc, dec
// 若当前同学在递增序列中，则当前同学的糖果比 pre 多 1 颗
// 若当前同学在递减序列中，则当前同学的糖果为 1 颗，并把当前递减序列中其他同学的糖果数量都加 1
// 若当前递减序列的长度等于上一个递增序列的长度
// 则要把上一个递增序列的最后一个同学也包含在当前递减序列中
func candy(ratings []int) int {
	inc, dec, pre, ans := 1, 0, 1, 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] > ratings[i-1] {
				pre++
			} else {
				pre = 1
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
