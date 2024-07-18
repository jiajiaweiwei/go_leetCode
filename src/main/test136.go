package main

func main() {

}

// 位运算 简单题 只出现一次的数字
// 方法1 使用异或运算(相同为0 不同为1) 且异或运算满足交换律和结合律
// 解题原理
func singleNumber(nums []int) int {
	resultNum := 0
	for _, v := range nums {
		resultNum ^= v
	}
	return resultNum
}

// 方法2 使用集合
func singleNumber2(nums []int) int {
	tempSet := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 如果集合中已经出现了该数字，就删除 没有就 添加
		if _, exist := tempSet[nums[i]]; exist {
			delete(tempSet, nums[i])
		} else {
			tempSet[nums[i]] = i
		}
	}
	// 剩下的唯一的数字
	for num := range tempSet {
		return num
	}
	return -1 // 如果没有找到返回一个标志值
}
