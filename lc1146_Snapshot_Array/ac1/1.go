package ac1

// 二分查找
// 在arr中记录每次set的id及对应的val值
// 并记录每次快照对应的set的id
// 最后在 arr[index] 中用二分查找找到snap_id对应setID的下标

type pair struct {
	setID int
	val   int
}

type SnapshotArray struct {
	setCount int
	snaps    []int
	arr      [][]pair
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{arr: make([][]pair, length)}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.setCount++
	this.arr[index] = append(this.arr[index], pair{setID: this.setCount, val: val})
}

func (this *SnapshotArray) Snap() int {
	this.snaps = append(this.snaps, this.setCount)
	return len(this.snaps) - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	setID := this.snaps[snap_id]
	l, r := 0, len(this.arr[index])
	for l < r {
		m := (r-l)>>1 + l
		if this.arr[index][m].setID > setID {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == 0 {
		return 0
	}
	return this.arr[index][l-1].val
}
