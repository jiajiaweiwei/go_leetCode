package main

func main() {
	reverseBits(9)

}

// 位运算 中等题 颠倒二进制位
// 方法1 位运算
func reverseBits(num uint32) uint32 {
	// 先将要移动的位置赋值，再移动
	// n & 1 n & 000000000000001(31个0) 按位与，同为1 才为1 取出最低位0 或 1
	// << (31 - i)  左移 将最低位 移动到 要反转的位置上
	// 通过按位或 将该位加到 返回值上
	// 返回值
	var resultNum uint32
	// 循环遍历  先获取 指定位置的值， 再 根据移位 操作 将指定位置的值放到返回值指定位置
	for i := 0; i < 31 && num > 0; i++ {
		// 计算顺序： 从右往左
		resultNum |= num & 1 << (31 - i)
		num >>= 1
	}
	return resultNum
}
