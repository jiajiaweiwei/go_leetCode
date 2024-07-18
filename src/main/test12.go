package main

func main() {

}

// 整数转罗马数字中等题 方法一：模拟
var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	//记录字符串 从最大罗马字符开始记录
	roman := []byte{}
	//遍历valueSymbols 从最大的字符开始
	for _, vs := range valueSymbols {
		//如果num够大
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}
