package ac1

// 归并排序
func sortArray(nums []int) []int {
	merge := func(a, b []int) []int {
		if len(a) == 0 {
			return b
		}
		if len(b) == 0 {
			return a
		}
		var res []int
		for i, j := 0, 0; i < len(a) || j < len(b); {
			if i >= len(a) {
				res = append(res, b[j])
				j++
			} else if j >= len(b) {
				res = append(res, a[i])
				i++
			} else if a[i] > b[j] {
				res = append(res, b[j])
				j++
			} else {
				res = append(res, a[i])
				i++
			}
		}
		return res
	}
	var sort func(l, r int) []int
	sort = func(l, r int) []int {
		if l == r-1 {
			return nums[l:r]
		}
		if l == r {
			return nil
		}
		m := (r-l)>>1 + l
		return merge(sort(l, m), sort(m, r))
	}
	return sort(0, len(nums))
}
