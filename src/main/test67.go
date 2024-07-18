package main

func main() {
	a, b := "11", "1"
	addBinary(a, b)
}

// 位运算 简单题 二进制求和
// 方法1 使用strconv.Atoi(a)将string转化为int ;错误 会导致溢出，应该用模拟位运算的方法
// 方法2 模拟位运算 这里可以用ascii码简化bitDemo方法
func addBinary(a string, b string) string {
	// 模拟位运算
	// 长度
	l1, l2 := len(a), len(b)
	// 获取最大长度
	maxL := l1
	if l1 < l2 {
		maxL = l2
	}
	// 返回值
	resultStrings := make([]byte, maxL+1)
	var goOn, this byte
	goOn = '0'
	i, j, k := l1-1, l2-1, maxL
	// 处理两个字符串都存在的部分
	for i >= 0 && j >= 0 {
		goOn, this = bitDemo(a[i], b[j], goOn)
		resultStrings[k] = this
		i--
		j--
		k--
	}
	// 处理剩余的部分
	for i >= 0 {
		goOn, this = bitDemo(a[i], '0', goOn)
		resultStrings[k] = this
		i--
		k--
	}
	for j >= 0 {
		goOn, this = bitDemo('0', b[j], goOn)
		resultStrings[k] = this
		j--
		k--
	}
	// 处理最高位的进位
	if goOn == '1' {
		resultStrings[0] = goOn
		return string(resultStrings)
	}
	return string(resultStrings[1:])
}

// 模拟位运算工具   根据 上一位进位，本两位  获取 下一位进位 和 本位
func bitDemo(b1, b2, goOnBefore byte) (goOn, this byte) {
	if goOnBefore == '1' {
		if b1 == '1' && b2 == '1' {
			return '1', '1'
		}
		if b1 == '0' && b2 == '0' {
			return '0', '1'
		}
		return '1', '0'
	}
	if goOnBefore == '0' {
		if b1 == '1' && b2 == '1' {
			return '1', '0'
		}
		if b1 == '0' && b2 == '0' {
			return '0', '0'
		}
		return '0', '1'
	}
	panic("传入值异常")
	return '0', '0'
}

// 方法2 使用ASCII码简化 求值
func addBinary2(a string, b string) string {
	// 长度
	l1, l2 := len(a), len(b)
	// 最大长度 与最大string
	maxL := l1
	if l2 > l1 {
		maxL = l2
	}
	// 返回值
	resultStrings := make([]byte, maxL+1)
	// 进位
	var carry byte
	// 下标
	i, j, k := l1-1, l2-1, maxL
	// 切片下标不溢出时
	for i >= 0 || j >= 0 {
		var bitA, bitB byte
		// a下标不溢出时
		if i >= 0 {
			// 字符 '0' 对应的 ASCII 值是 48，而字符 '1' 对应的 ASCII 值是 49
			bitA = a[i] - '0' // convert to byte 0 or 1
			i--
		} else {
			bitA = 0
		}
		// b下标不溢出时
		if j >= 0 {
			bitB = b[j] - '0' // convert to byte 0 or 1
			j--
		} else {
			bitB = 0
		}
		// 三位求总和
		sum := bitA + bitB + carry
		// 赋值 ’0‘ 或 ’1‘ 的ascii码
		resultStrings[k] = sum%2 + '0' // convert sum back to byte '0' or '1'
		// 判断进位
		carry = sum / 2
		k--
	}
	// 最后一个进位
	if carry == 1 {
		resultStrings[0] = '1'
		return string(resultStrings)
	}
	return string(resultStrings[1:])
}
