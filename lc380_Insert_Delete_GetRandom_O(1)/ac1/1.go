package ac1

import "math/rand/v2"

// 切片用于 O(1) 随机读取，哈希表用于 O(1) 插入、删除

type RandomizedSet struct {
	set map[int]int
	sli []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
		sli: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.sli = append(this.sli, val)
	this.set[val] = len(this.sli) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}
	this.sli[index] = this.sli[len(this.sli)-1]
	this.set[this.sli[index]] = index
	this.sli = this.sli[:len(this.sli)-1]
	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.sli[rand.N(len(this.sli))]
}
