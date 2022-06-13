package ac1

type TopVotedCandidate struct {
	winners []int
	times   []int
}

// 预处理，先计算出投了第i张选票时的获胜者
func Constructor(persons []int, times []int) TopVotedCandidate {
	count := make([]int, len(persons))
	max := 0
	winner := 0
	for i := 0; i < len(persons); i++ {
		count[persons[i]]++
		if count[persons[i]] >= max {
			max = count[persons[i]]
			winner = persons[i]
		}
		persons[i] = winner
	}
	return TopVotedCandidate{winners: persons, times: times}
}

// 二分搜索，找出时刻t投到第l张选票（upper_bound），则winner[l-1]即为答案
func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.times)
	for l < r {
		m := (r-l)>>1 + l
		if this.times[m] > t {
			r = m
		} else {
			l = m + 1
		}
	}
	return this.winners[l-1]
}
