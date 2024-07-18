package main

func main() {

}

// 位运算 中等题 数字范围按位与 （暴力法不可取，题目条件不允许暴力）
//
//	经过数学推断 我们可以将问题重新表述为：给定两个整数，我们要找到它们对应的二进制字符串的公共前缀
func rangeBitwiseAnd(left int, right int) int {
	temp := 0
	for left < right {
		left, right = left>>1, right>>1
		temp++
	}
	return left << temp
}
