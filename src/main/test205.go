package main

import (
	"fmt"
	"strings"
)

func main() {
	// 要切割的字符串
	str := "Hello World Go"

	// 通过空格切割字符串
	parts := strings.Split(str, " ")

	// 打印切割后的结果
	for i, part := range parts {
		fmt.Printf("Part %d: %s\n", i, part)
	}
}

// 哈希表 简单题 同构字符串
// 方法1  使用哈希表 存s的字符集 再判断t中有无映射不正确的字符集
func isIsomorphic(s string, t string) bool {
	// 特殊情况处理 长度不一样直接返回false
	if len(s) != len(t) {
		return false
	}
	// 哈希表 由于需要不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上 故需要两个哈希表
	mapping := make(map[byte]byte)
	reverseMapping := make(map[byte]byte)
	//s t存入哈希表
	i := 0
	for i < len(s) {
		if val, ok := mapping[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			mapping[s[i]] = t[i]
		}

		if val, ok := reverseMapping[t[i]]; ok {
			if val != s[i] {
				return false
			}
		} else {
			reverseMapping[t[i]] = s[i]
		}
		i++
	}
	return true
}
