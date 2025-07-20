package ac1

import (
	"strconv"
	"strings"
)

// 字符串分割
func compareVersion(version1 string, version2 string) int {
	version1List := strings.Split(version1, ".")
	version2List := strings.Split(version2, ".")
	for i := 0; i < len(version1List) || i < len(version2List); i++ {
		v1, v2 := 0, 0
		if i < len(version1List) {
			v1, _ = strconv.Atoi(version1List[i])
		}
		if i < len(version2List) {
			v2, _ = strconv.Atoi(version2List[i])
		}
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	return 0
}
