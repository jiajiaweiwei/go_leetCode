package main

import "strings"

func main() {

}

// 哈希表 简单题 单词规律
// 方法1 使用strings.Split接口切割字符串 存入哈希表对比
func wordPattern(pattern string, s string) bool {
	// 通过空格切割字符串
	parts := strings.Split(s, " ")
	// 特殊情况处理
	if len(pattern) != len(parts) {
		return false
	}
	// 存入哈希表
	hashMap1 := make(map[byte]string)
	hashMap2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if val, ok := hashMap1[pattern[i]]; ok {
			if val != parts[i] {
				return false
			}
		} else {
			hashMap1[pattern[i]] = parts[i]
		}

		if val, ok := hashMap2[parts[i]]; ok {
			if val != pattern[i] {
				return false
			}
		} else {
			hashMap2[parts[i]] = pattern[i]
		}
	}
	return true
}
