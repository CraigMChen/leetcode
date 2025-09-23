package ac1

import "unsafe"

// 递归
// decode 函数将形如 N[xxx] 的字符串解码
// 遇到嵌套的括号就求出该括号的重复次数 N 和括号的起点和终点，递归调用
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func decodeString(s string) string {
	return bytes2string(decode(1, string2bytes(s)))
}

func decode(time int, s []byte) []byte {
	var res []byte
	for i := 0; i < len(s); {
		if s[i] >= 'a' && s[i] <= 'z' {
			res = append(res, s[i])
			i++
			continue
		}

		t := 0
		j := i
		for ; j < len(s) && s[j] != '['; j++ {
			t = t*10 + int(s[j]-'0')
		}
		j++
		start := j
		count := -1
		for ; j < len(s) && count != 0; j++ {
			if s[j] == '[' {
				count--
			} else if s[j] == ']' {
				count++
			}
		}
		res = append(res, decode(t, s[start:j])...)
		i = j
	}

	var ans []byte
	for i := 0; i < time; i++ {
		ans = append(ans, res...)
	}
	return ans
}

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
