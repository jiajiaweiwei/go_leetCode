package main

import "sort"

func main() {
	s := "aacc"
	t := "ccac"
	isAnagram(s, t)
}

// 哈希表 简单题 有效的字母异位词
// 方法1 使用哈希表 存各个字符的出现次数
func isAnagram(s string, t string) bool {
	// 特殊情况处理
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[byte]int)
	// 存储 s字符出现次数
	for i := 0; i < len(s); i++ {
		if _, ok := hashMap[s[i]]; ok {
			hashMap[s[i]]++
		} else {
			hashMap[s[i]] = 1
		}
	}
	// 判断 t是否是有效的
	for i := 0; i < len(t); i++ {
		if _, ok := hashMap[t[i]]; ok {
			hashMap[t[i]]--
			if hashMap[t[i]] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// 方法2 先排序，后判断 时间复杂度更低
func isAnagram2(s string, t string) bool {
	// 特殊情况处理
	if len(s) != len(t) {
		return false
	}
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}
