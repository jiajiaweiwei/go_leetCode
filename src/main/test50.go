package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(3 % 2)
}

// 数学 中等题 实现幂函数的计算
// 方法1 半暴力法 超时
func myPow(x float64, n int) float64 {
	if x > 0 {
		return math.Exp(float64(n) * math.Log(float64(x)))
	} else if x == 0 {
		return 0
	} else {
		// 对负数求幂函数
		resultNum := 1.0
		if n >= 0 {
			for i := 0; i < n; i++ {
				resultNum = resultNum * x
			}
			return resultNum
		} else {
			for i := 0; i < -n; i++ {
				resultNum = resultNum * x
			}
			return 1 / resultNum
		}

	}
}

// 方法2 快速幂 + 递归
func myPow2(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

// 快速幂方法
func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	// 递归 类似二分法的思想
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
