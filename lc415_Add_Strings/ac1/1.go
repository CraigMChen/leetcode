package ac1

// 模拟
func addStrings(num1 string, num2 string) string {
	var (
		res   []byte
		add   string
		carry byte = 0
	)
	if len(num1) > len(num2) {
		res = []byte(num1)
		add = num2
	} else {
		res = []byte(num2)
		add = num1
	}

	for i, j := len(res)-1, len(add)-1; i >= 0; i-- {
		var a, b byte = res[i], '0'
		if j >= 0 {
			b = add[j]
		}
		carry, res[i] = byteAdd(carry, a, b)
		j--
	}
	if carry > 0 {
		return "1" + string(res)
	}
	return string(res)
}

func byteAdd(carry, a, b byte) (byte, byte) {
	sum := a - '0' + b - '0' + carry
	carry = sum / 10
	return carry, sum%10 + '0'
}
