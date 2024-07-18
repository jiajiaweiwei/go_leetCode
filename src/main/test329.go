package main

import "fmt"

func main() {
	s := "acb"
	subsequence := isSubsequence(s, "ahbgdc")
	fmt.Println(subsequence)
}

// 双指针简单题 判断子序列
func isSubsequence(s string, t string) bool {
	//特殊情况处理
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) || len(t) == 0 {
		return false
	}
	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		for s[i] != t[j] {
			j++
			if j == len(t) {
				return false
			}
		}
		i++
	}
	return i == len(s)-1
}
