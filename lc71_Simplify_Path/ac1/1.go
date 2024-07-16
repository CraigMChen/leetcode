package ac1

import "strings"

// 栈
func simplifyPath(path string) string {
	var stack []string
	strs := strings.Split(path, "/")
	for _, str := range strs {
		if str == "" || str == "." {
			continue
		}
		if str == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, str)
		}
	}

	ans := "/" + strings.Join(stack, "/")
	return ans
}
