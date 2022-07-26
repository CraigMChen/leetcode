package ac2

type TopVotedCandidate struct {
	votes map[int][]int
}

// 记录每个人获得选票的时间
func Constructor(persons []int, times []int) TopVotedCandidate {
	votes := make(map[int][]int)
	for i := 0; i < len(persons); i++ {
		votes[persons[i]] = append(votes[persons[i]], times[i])
	}
	return TopVotedCandidate{votes: votes}
}

// 二分查找
func (this *TopVotedCandidate) Q(t int) int {
	winner := -1
	last := -1
	max := -1
	for person, votes := range this.votes {
		l, r := 0, len(votes)
		for l < r {
			m := (r-l)>>1 + l
			if votes[m] > t {
				r = m
			} else {
				l = m + 1
			}
		}
		if l > 0 && (l > max || l == max && votes[l-1] > last) {
			winner = person
			max = l
			last = votes[l-1]
		}
	}
	return winner
}
