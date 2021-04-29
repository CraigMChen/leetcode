package ac1

type TopVotedCandidate struct {
	winner []int
	times []int
}

// 预处理，先计算出投了第i张选票时的获胜者
func Constructor(persons []int, times []int) TopVotedCandidate {
	max := -1
	p := -1
	counter := make(map[int]int)
	winner := make([]int, len(persons))
	for i := 0; i < len(persons); i++ {
		counter[persons[i]]++
		if counter[persons[i]] >= max {
			max = counter[persons[i]]
			p = persons[i]
		}
		winner[i] = p
	}
	return TopVotedCandidate{
		winner: winner,
		times:  times,
	}
}

// 二分搜索，找出时刻t投到第l张选票（upper_bound），则winner[l-1]即为答案
func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.times)
	for l < r {
		m := (r - l) >> 1 + l
		if this.times[m] <= t {
			l = m + 1
		} else {
			r = m
		}
	}
	return this.winner[l - 1]
}
