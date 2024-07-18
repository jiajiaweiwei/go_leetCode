package main

func main() {

}

// 双指针 中等题 盛水最多的容器
// 方法1 双指针法 每次移动较短的那个柱子
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	resultValue := 0 //面积
	h := 0
	for i < j {
		h = min(height[i], height[j])
		resultValue = max(resultValue, (j-i)*h)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return resultValue
}

// 求最大值
func max11(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 求最小值
func min11(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
