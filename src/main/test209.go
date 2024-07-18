package main

func main() {

}

// 滑动窗口中等题 长度最小的子数组
// 方法1 使用滑动窗口算法 左边界和右边界只向右移，右边界无需左移，因为右边界左移一定不满足
func minSubArrayLen(target int, nums []int) int {
	// 特殊情况处理
	if nums[0] >= target {
		return 1
	}

	// 记录满足条件的长度，取最小
	start, end := 0, 0
	tempSum := 0
	resultValue := len(nums) + 1

	for start < len(nums) {
		if tempSum < target && end < len(nums) {
			tempSum += nums[end]
			end++
		} else {
			if tempSum >= target {
				resultValue = min209(resultValue, end-start)
			}
			tempSum -= nums[start]
			start++
		}
	}

	if resultValue == len(nums)+1 {
		return 0
	}
	return resultValue
}

// 求最小值
func min209(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
