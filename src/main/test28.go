package main

import "fmt"

func main() {

	ter()

}
func ter() {
	for _, q := range "123" {
		fmt.Println(q)
	}
}

// 力扣28简单题 字符串的模式匹配

// 方法1 暴力法
func strStr(haystack string, needle string) int {
	//特殊情况处理
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		j := 0
		temp := i
		for j < len(needle) {
			if haystack[i] == needle[j] {
				i++
				j++
			} else {
				goto start
			}
			if j == len(needle) {
				return i - len(needle)
			}
		}
	start:
		i = temp

	}
	return -1
}

// 方法1 暴力法
func strStr1(haystack, needle string) int {
	n, m := len(haystack), len(needle)
outer: //用来标记外层循环，此时使用continue不会影响i的值
	//遍历次数 =  两字符串长度之差
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] { //判断字符是否合规
				continue outer //不合规就直接跳出该层循环 为了跳过return i 进入下一轮循环
			}
		}
		return i
	}
	return -1
}

// 方法2 KMP算法 求next数组
func strStr2(haystack, needle string) int {
	//KMP算法
	//求出next数组
	next := getNextVal(needle)
	//根据next数组进行字符串匹配
	n, m := len(haystack), len(needle)
	next = getNextVal(needle)
	j := 0
	for i := 0; i < n; i++ {
		for j != -1 && haystack[i] != needle[j] {
			j = next[j]
		}
		if j == m-1 {
			return i - m + 1
		}
		j++
	}

	return -1
}

// 重点 使用递归求next数组 （本质 寻找子串中相同前后缀的长度）
func getNextVal(needle string) []int {
	n := len(needle)            //needle字符串长度
	next := make([]int, n)      //要求的next数组
	next[0] = -1                //初始值确定
	var calcNext func(i, k int) //递归函数
	calcNext = func(i, k int) { //i表示next数组下标 k表示上一个递归数组元素值
		if i == n { //递归结束信号
			return
		}
		if k == -1 || needle[i] == needle[k] {
			next[i] = k + 1    // 如果前缀和相同 递归数组值等于上一个递归数组元素值+1
			calcNext(i+1, k+1) //递归下一个位置
		} else { //前缀和不同，传递next[k]的值
			calcNext(i, next[k]) //递归
		}
	}
	calcNext(1, 0) //递归 从下标0开始
	return next
}
