package main

import (
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 7, 8, 9, 10, 11, 12, 13, 14}
	longestConsecutive1(nums)
}

// 哈希表 中等题 最长连续序列 要求使用时间复杂度为On 比排序的On*logn好
// 方法1 先排序后便利
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	// 返回值
	longestStreak := 1
	currentStreak := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { // 跳过重复元素
			if nums[i] == nums[i-1]+1 {
				currentStreak++
			} else {
				if currentStreak > longestStreak {
					longestStreak = currentStreak
				}
				currentStreak = 1
			}
		}
	}
	// 最后一次 streak 可能是最长的
	if currentStreak > longestStreak {
		longestStreak = currentStreak
	}
	return longestStreak
}

// 方法2 哈希表
func longestConsecutive1(nums []int) int {
	// 使用哈希表
	// 先将所有数加入到哈希表中
	hashMap := make(map[int]bool)
	for _, v := range nums {
		hashMap[v] = true
	}
	// 返回值
	resultNum := 0
	// 临时变量
	temp := 0
	// 判断最长子序列
	for v := range hashMap {
		if !hashMap[v-1] {
			// 只判断是序列头部的元素
			temp = 1
			cur := v
			for hashMap[cur+1] {
				cur++
				temp++
			}
			resultNum = max128(resultNum, temp)
		}

	}
	return resultNum
}
func max128(a, b int) int {
	if a > b {
		return a
	}
	return b
}
