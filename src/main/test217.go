package main

import "sort"

func main() {

}

// 哈希表 简单题 存在重复元素1  O(n)O(n)
// 方法1 使用哈希表
func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int)
	// 循环存入哈希表
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; !ok {
			hashMap[nums[i]] = i
		} else {
			return true
		}
	}
	return false
}

// 方法2 排序 优点 时间复杂度比 方法1 低 O(logn)O(logn)
func containsDuplicate2(nums []int) bool {
	// 直接将nums数组排序
	sort.Ints(nums)
	//遍历查看是否有重复元素
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
