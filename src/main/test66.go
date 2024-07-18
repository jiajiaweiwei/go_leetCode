package main

func main() {

}

// 数学 简单题 加1
// 方法1 找出末尾有多少个9即可
func plusOne(digits []int) []int {
	l := len(digits)
	// 找出末尾连续9的个数
	temp := l - 1
	num := 0
	// 检查是否越界并统计连续9的个数
	for temp >= 0 && digits[temp] == 9 {
		temp--
		num++
	}
	// 如果所有数字都是9，需要扩展数组长度
	if num == l {
		result := make([]int, l+1)
		result[0] = 1
		return result
	}
	// 增加非9位置的值
	digits[l-num-1]++
	// 将末尾连续9的位置设为0
	for i := l - 1; i > l-num-1; i-- {
		digits[i] = 0
	}
	return digits
}
