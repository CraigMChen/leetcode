package ac1

type TimeMap struct {
	items map[string][]item
}

type item struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{items: make(map[string][]item)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.items[key] = append(this.items[key], item{
		timestamp: timestamp,
		value:     value,
	})
}

// 二分查找
func (this *TimeMap) Get(key string, timestamp int) string {
	items := this.items[key]
	l, r := 0, len(items)
	for l < r {
		m := (r-l)>>1 + l
		if items[m].timestamp > timestamp {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == 0 {
		return ""
	}
	return items[l-1].value
}
