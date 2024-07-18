package main

func main() {

}

// 力扣58最后一个单词的长度
// 法1 反向遍历法
func lengthOfLastWord(s string) int {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	i := 0
	for s[index] != ' ' {
		index--
		i++
	}
	return i
}
