package main

func main() {

}

// 数学 中等题 阶乘之后的0
// 解题技巧：
// 含有 2 的因子每两个出现一次
// 含有 5 的因子每 5 个出现一次，
// 所有 2 出现的个数远远多于 5，
// 换言之找到一个 5，一定能找到一个 2 与之配对。
// 所以我们只需要找有多少个 5（找到有多少个数的因子有5）
func trailingZeroes(n int) int {
	resultNum := 0
	for i := n; i > 0; i-- {
		// 子循环找因子
		temp := i
		for temp > 0 {
			if temp%5 == 0 { // 一开始就不是5的倍数就可以直接退出了
				resultNum++
				temp /= 5
			} else {
				break
			}
		}
	}
	return resultNum
}
