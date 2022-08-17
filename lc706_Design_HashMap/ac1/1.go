package ac1

import "container/list"

// 链地址法

const BASE = 877

type pair struct {
	key   int
	value int
}

type MyHashMap struct {
	list []list.List
}

func hash(x int) int {
	return x % BASE
}

func Constructor() MyHashMap {
	return MyHashMap{list: make([]list.List, BASE, BASE)}
}

func (this *MyHashMap) Put(key int, value int) {
	h := hash(key)
	for i := this.list[h].Front(); i != nil; i = i.Next() {
		p := i.Value.(pair)
		if p.key == key {
			p.value = value
			i.Value = p
			return
		}
	}
	this.list[h].PushBack(pair{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := hash(key)
	for i := this.list[h].Front(); i != nil; i = i.Next() {
		p := i.Value.(pair)
		if p.key == key {
			return p.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := hash(key)
	for i := this.list[h].Front(); i != nil; i = i.Next() {
		p := i.Value.(pair)
		if p.key == key {
			this.list[h].Remove(i)
		}
	}
}
