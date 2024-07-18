package main

func main() {

}

// 位运算 中等题 只出现一次的数字2 要求是线性的时间复杂度 和 常量级的空间复杂度
// 方法1 原始方法 哈希表
func singleNumber4(nums []int) int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]]++
	}
	for i, v := range hashMap {
		if v == 1 {
			return i
		}
	}
	return 0
}

// 方法2 使用位运算 构造逻辑电路（重点）
func singleNumber3(nums []int) int {
	resultNum := 0
	// 一个数字出现一次 其他数字都是出现三次
	for i := 0; i < len(nums); i++ {

	}

	return resultNum
}
