package main

import "unsafe"

// 模拟竖式计算
//func multiply(num1 string, num2 string) string {
//	if num1 == "0" || num2 == "0" {
//		return "0"
//	}
//	addWithCarry := func(carry, a, b byte) (byte, byte) {
//		sum := a - '0' + b - '0' + carry
//		carry = sum / 10
//		return carry, sum%10 + '0'
//	}
//	mutWithCarry := func(carry, a, b byte) (byte, byte) {
//		product := (a-'0')*(b-'0') + carry
//		carry = product / 10
//		return carry, product%10 + '0'
//	}
//	strAdd := func(n1, n2 []byte) []byte {
//		var carry byte = 0
//		sum, add := n1, n2
//		if len(n1) < len(n2) {
//			sum, add = n2, n1
//		}
//		for i, j := len(sum)-1, len(add)-1; i >= 0; i-- {
//			if j >= 0 {
//				carry, sum[i] = addWithCarry(carry, sum[i], add[j])
//			} else {
//				carry, sum[i] = addWithCarry(carry, sum[i], '0')
//			}
//			j--
//		}
//		if carry > 0 {
//			sum = append([]byte{'1'}, sum...)
//		}
//		return sum
//	}
//	strMut := func(x byte, y []byte) []byte {
//		var carry byte = 0
//		product := y
//		for i := len(product) - 1; i >= 0; i-- {
//			carry, product[i] = mutWithCarry(carry, x, product[i])
//		}
//		if carry > 0 {
//			product = append([]byte{carry + '0'}, product...)
//		}
//		return product
//	}
//
//	var (
//		carry   []byte
//		product = []byte{'0'}
//	)
//	for i := len(num1) - 1; i >= 0; i-- {
//		p := append(strMut(num1[i], []byte(num2)), carry...)
//		product = strAdd(product, p)
//		carry = append(carry, '0')
//	}
//	return string(product)
//}

func multiply(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	ans := "0"
	for i := 0; i < len(num1); i++ {
		ans = add(ans, digitMultiply(num1[i], num2, len(num1)-i-1))
	}
	return ans
}

func add(num1 string, num2 string) string {
	var ans []byte
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum := x + y + carry
		ans = append(ans, byte(sum%10)+'0')
		carry = sum / 10
	}
	if carry > 0 {
		ans = append(ans, byte(carry)+'0')
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return *(*string)(unsafe.Pointer(&ans))
}

func digitMultiply(digit byte, num string, left int) string {
	if digit == '0' {
		return "0"
	}
	var ans []byte
	carry := 0
	for i := len(num) - 1; i >= 0; i-- {
		product := int(digit-'0')*int(num[i]-'0') + carry
		ans = append(ans, byte(product%10)+'0')
		carry = product / 10
	}
	if carry > 0 {
		ans = append(ans, byte(carry)+'0')
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	for i := 0; i < left; i++ {
		ans = append(ans, '0')
	}
	return *(*string)(unsafe.Pointer(&ans))
}
