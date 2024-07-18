package main

import "fmt"

func main() {
	s := "0P"
	palindrome := isPalindrome(s)
	fmt.Println(palindrome)
	b := "AZaz09"
	for i, _ := range b {
		fmt.Print(b[i], "-")
	}
}

// 验证回文串
// 方法1 使用双指针+ASCII码
func isPalindrome(s string) bool {
	l := len(s) - 1
	i, j := 0, l
	for i < j && i <= l && i >= 0 && j <= l && j >= 0 {
		//判断是不是字符，是的话左右指针对比，不是就跳过
		i, j = jump(i, j, s)
		if i == -1 {
			return true
		}
		if s[i] != s[j] {
			//判断是否是相同字母（数字和字符也可能相差32 会误判）
			//去除数字误判成字母
			//有其中之一是数字
			if (s[i] >= 48 && s[i] <= 57) || (s[j] >= 48 && s[j] <= 57) {
				return false
			}
			if s[i] > s[j] && (s[i]-s[j]) != 32 {
				return false
			}
			if s[i] < s[j] && (s[j]-s[i]) != 32 {
				return false
			}
		}
		i++
		j--
	}
	return true
}

// 跳过不是字符的位置
func jump(a, b int, s string) (i, j int) {
	//不是字符
	for !((s[a] >= 'A' && s[a] <= 'Z') || (s[a] >= 'a' && s[a] <= 'z') || (s[a] >= '0' && s[a] <= '9')) {
		a++
		if a > len(s)-1 {
			return -1, -1
		}
	}
	for !((s[b] >= 'A' && s[b] <= 'Z') || (s[b] >= 'a' && s[b] <= 'z') || (s[b] >= '0' && s[b] <= '9')) {
		b--
		if b < 0 {
			return -1, -1
		}
	}
	return a, b
}
