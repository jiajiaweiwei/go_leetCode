package main

import "math"

func main() {

}

// 数学 简单题 x的平方根(向下取整)
// 方法1 袖珍计算器算法（使用指数函数和对数函数得到算术平方根） math.Log求对数 math.Exp求指数
func mySqrt(x int) int {
	// golang 原生库提供的接口math.Sqrt()
	// 特殊情况处理
	if x == 0 {
		return 0
	}
	resultNum := int(math.Exp(0.5 * math.Log(float64(x))))
	if (resultNum+1)^2 > x {
		return resultNum
	}
	return resultNum + 1
}

// 方法2 二分查找
// 方法3 牛顿迭代 （也是golang原生math.Sqrt()的实现原理）
