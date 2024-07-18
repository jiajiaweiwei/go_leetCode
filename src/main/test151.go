package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	// 去除字符串两端空格
	s = strings.TrimSpace(s)
	// 按空格分割字符串为单词数组
	words := strings.Fields(s)
	// 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// 拼接单词数组为字符串
	return strings.Join(words, " ")
}
