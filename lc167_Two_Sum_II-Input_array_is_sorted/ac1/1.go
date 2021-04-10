package ac1

// 二分
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		l, r := i + 1, len(numbers)
		for l < r {
			m := (r - l) / 2 + l
			if numbers[m] + numbers[i] == target {
				return []int{i + 1, m + 1}
			} else if numbers[m] + numbers[i] < target {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	return nil
}
