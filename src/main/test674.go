package main

func main() {

}

// 贪心 简单题 求连续最长子序列
// 方法1 暴力法 直接遍历 时间复杂度On
func findLengthOfLCIS(nums []int) int {
	// 返回值
	resultNum := 0
	for i := 0; i < len(nums); i++ {
		tempNum := 1
		for i+1 < len(nums) && nums[i] < nums[i+1] {
			tempNum++
			i++
		}
		resultNum = max674(resultNum, tempNum)
	}
	return resultNum
}
func max674(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法2 贪心 时间复杂度也为On
// 为了得到最长连续递增序列，可以使用贪心的策略得到尽可能长的连续递增序列。
// 做法是
// 使用记录当前连续递增序列的开始下标和结束下标，
// 遍历数组的过程中每次比较相邻元素，
// 根据相邻元素的大小关系决定是否需要更新连续递增序列的开始下标
func findLengthOfLCIS2(nums []int) int {
	// 返回值
	resultNum := 0
	start := 0
	// 从局部最优到全局最优
	for i, v := range nums {
		// 不满足递增条件就改变起点
		if i > 0 && nums[i-1] >= v { //连续递增 非严格递增
			start = i
		}
		// 满足条件就更新最长子序列的值
		resultNum = max674(resultNum, i-start+1)
	}
	return resultNum
}
