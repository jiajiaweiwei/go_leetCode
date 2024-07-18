package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(i)) // int在 64位处理器上 占8字节 32位置上占4字节
}

// 哈希表 简单题 判断快乐数
// 什么是 快乐数
/*快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false*/

// 方法1 哈希表
// 使用哈希表判断 已经计算出的结果是否已经出现过，出现过就会导致循环，从而不是快乐数字
func isHappy(n int) bool {
	// 构建哈希表
	hashMap := make(map[int]bool)
	for ; n != 1 && !hashMap[n]; n, hashMap[n] = step(n), true {
	}
	return n == 1
}

// 计算平方和辅助函数
func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// 方法2 使用双指针
