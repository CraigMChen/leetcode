package ac1

// 分治
// 时间复杂度 O(nlogn)
// 空间复杂度 O(logn)
func majorityElement(nums []int) int {
	var getMaj func(int, int) int
	getMaj = func(l, r int) int {
		if l >= r {
			return -1
		}
		if r-l == 1 {
			return nums[l]
		}
		m := (r-l)>>1 + l
		majL, majR := getMaj(l, m), getMaj(m, r)
		if majL == majR {
			return majL
		}
		countL, countR := 0, 0
		for i := l; i < r; i++ {
			if nums[i] == majL {
				countL++
			}
			if nums[i] == majR {
				countR++
			}
		}
		if countL > countR {
			return majL
		} else {
			return majR
		}
	}
	return getMaj(0, len(nums))
}
