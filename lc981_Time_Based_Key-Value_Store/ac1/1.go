package ac1

// value[key][i]表示键key在第i次插入操作之后的值
// timestamp[key[i]表示键key的第i次插入操作的时间
type TimeMap struct {
	value     map[string][]string
	timestamp map[string][]int
}

func Constructor() TimeMap {
	return TimeMap{
		value:     make(map[string][]string),
		timestamp: make(map[string][]int),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.value[key] == nil {
		this.value[key] = make([]string, 1)
		this.timestamp[key] = make([]int, 1)
	}
	this.value[key] = append(this.value[key], value)
	this.timestamp[key] = append(this.timestamp[key], timestamp)
}

// 从key对应的时间数组中用二分搜索找出timestamp对应的是哪一次操作
func (this *TimeMap) Get(key string, timestamp int) string {
	if this.value[key] == nil {
		return ""
	}
	l, r := 0, len(this.timestamp[key])
	for l < r {
		m := (r - l) >> 1 + l
		if this.timestamp[key][m] <= timestamp {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == 0 {
		return ""
	}
	return this.value[key][l - 1]
}
