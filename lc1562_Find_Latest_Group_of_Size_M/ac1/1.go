package ac1

// 模拟
func findLatestStep(arr []int, m int) int {
	var (
		count = 0
		res   = -1
	)
	endpoints := make([][2]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		left, right := arr[i], arr[i]
		if arr[i] > 1 && endpoints[arr[i]-1][0] != 0 {
			left = endpoints[arr[i]-1][0]
			leftLength := endpoints[arr[i]-1][1] - endpoints[arr[i]-1][0] + 1
			if leftLength == m {
				count--
			}
		}
		if arr[i] < len(arr) && endpoints[arr[i]+1][1] != 0 {
			right = endpoints[arr[i]+1][1]
			rightLength := endpoints[arr[i]+1][1] - endpoints[arr[i]+1][0] + 1
			if rightLength == m {
				count--
			}
		}
		length := right - left + 1
		if length == m {
			count++
		}
		if count > 0 {
			res = i + 1
		}
		endpoints[left], endpoints[right] = [2]int{left, right}, [2]int{left, right}
	}
	return res
}
