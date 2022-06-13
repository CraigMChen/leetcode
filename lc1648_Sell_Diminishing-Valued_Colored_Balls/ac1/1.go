package ac1

const MOD int = 1e9 + 7

// 二分查找
// 可以先将所有的球都取到只剩m个（inventory[i]>m）
// 此时取出的总数量为count
// 若count < orders，则再分别从 order - count 堆含有m个的球堆中取1个
// 可以通过二分查找来确定m
func maxProfit(inventory []int, orders int) int {
	l, r := 0, 0
	for i := 0; i < len(inventory); i++ {
		if inventory[i]+1 > r {
			r = inventory[i] + 1
		}
	}

	getCount := func(m int) int {
		count := 0
		for i := 0; i < len(inventory); i++ {
			if inventory[i] > m {
				count += inventory[i] - m
			}
		}
		return count
	}
	for l < r {
		m := (r-l)>>1 + l
		if getCount(m) <= orders {
			r = m
		} else {
			l = m + 1
		}
	}

	var sum, count = 0, 0
	for i := 0; i < len(inventory); i++ {
		if inventory[i] > l {
			count += inventory[i] - l
			sum = (sum%MOD + (inventory[i]+l+1)*(inventory[i]-l)/2%MOD) % MOD
		}
	}
	if count < orders {
		sum = (sum%MOD + (orders-count)*l%MOD) % MOD
	}
	return sum
}
