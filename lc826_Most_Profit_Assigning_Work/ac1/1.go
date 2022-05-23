package ac1

import "sort"

// 双指针
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	type pair struct {
		difficulty int
		profit     int
	}
	jobs := make([]pair, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		jobs[i] = pair{
			difficulty: difficulty[i],
			profit:     profit[i],
		}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].difficulty < jobs[j].difficulty
	})
	sort.Ints(worker)

	res := 0
	best := 0
	for i, j := 0, 0; j < len(worker); j++ {
		for i < len(difficulty) {
			if jobs[i].difficulty > worker[j] {
				break
			}
			best = max(best, jobs[i].profit)
			i++
		}
		res += best
	}
	return res
}
