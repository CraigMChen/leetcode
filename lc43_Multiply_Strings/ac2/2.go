package ac2

import "unsafe"

// 乘法
// num1 * num2 的结果最大长度为 len(num1)+len(num2)，最小长度为 len(num1)+len(num2)-1
// 用一个长度为 len(num1)+len(num2) 的数组 product 保存乘法结果
// num1[i] * num2[j] 的结果累加到 product[i+j+1] 上
// 然后将 product 数组中每一位进位都累加到前一位中
// 最后将 product 转换为字符串
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	product := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	for i := len(product) - 1; i > 0; i-- {
		product[i-1] += product[i] / 10
		product[i] = product[i] % 10
	}
	if product[0] == 0 {
		product = product[1:]
	}
	res := make([]byte, len(product))
	for i := 0; i < len(product); i++ {
		res[i] = byte(product[i]) + '0'
	}
	return *(*string)(unsafe.Pointer(&res))
}
