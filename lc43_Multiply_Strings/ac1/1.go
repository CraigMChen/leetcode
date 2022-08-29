package main

// 模拟竖式计算
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	addWithCarry := func(carry, a, b byte) (byte, byte) {
		sum := a - '0' + b - '0' + carry
		carry = sum / 10
		return carry, sum%10 + '0'
	}
	mutWithCarry := func(carry, a, b byte) (byte, byte) {
		product := (a-'0')*(b-'0') + carry
		carry = product / 10
		return carry, product%10 + '0'
	}
	strAdd := func(n1, n2 []byte) []byte {
		var carry byte = 0
		sum, add := n1, n2
		if len(n1) < len(n2) {
			sum, add = n2, n1
		}
		for i, j := len(sum)-1, len(add)-1; i >= 0; i-- {
			if j >= 0 {
				carry, sum[i] = addWithCarry(carry, sum[i], add[j])
			} else {
				carry, sum[i] = addWithCarry(carry, sum[i], '0')
			}
			j--
		}
		if carry > 0 {
			sum = append([]byte{'1'}, sum...)
		}
		return sum
	}
	strMut := func(x byte, y []byte) []byte {
		var carry byte = 0
		product := y
		for i := len(product) - 1; i >= 0; i-- {
			carry, product[i] = mutWithCarry(carry, x, product[i])
		}
		if carry > 0 {
			product = append([]byte{carry + '0'}, product...)
		}
		return product
	}

	var (
		carry   []byte
		product = []byte{'0'}
	)
	for i := len(num1) - 1; i >= 0; i-- {
		p := append(strMut(num1[i], []byte(num2)), carry...)
		product = strAdd(product, p)
		carry = append(carry, '0')
	}
	return string(product)
}
