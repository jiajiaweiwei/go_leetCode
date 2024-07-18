package main

import "fmt"

func main() {
	toInt := romanToInt("MCMXCIV")
	fmt.Println(toInt)
	//fmt.Println('I') // Go 语言中，单引号 ' ' 通常用于表示 Unicode 字符（rune）
}

// 力扣13 简单题 罗马数字转整数
// 方法1 一个个比较 左边字符比右边字符小，则减去右边字符，否则加
func romanToInt(s string) int {
	//设置一个map存放映射关系 byte无符号整数 ASCII码
	var defaultMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	//返回值
	firstValue := defaultMap[s[0]]
	result := firstValue
	for i := 1; i < len(s); i++ {
		//前一个
		if defaultMap[s[i]] > firstValue {
			result -= firstValue * 2
			result += defaultMap[s[i]]
		} else {
			result += defaultMap[s[i]]
		}
		firstValue = defaultMap[s[i]]
	}
	return result
}
