package main

func main() {

}

// 哈希表 简单题 两数之和
// 方法1 哈希表
func twoSum1(nums []int, target int) []int {
	// 将 nums存入哈希表
	hashMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]] = i
	}
	// 找 符合条件的两个数
	// 注意 数组中同一个元素在答案里不能重复出现。
	for i := 0; i < len(nums); i++ {
		goalNum := target - nums[i]
		if v, ok := hashMap[goalNum]; ok && i != v {
			return []int{i, v}
		}
	}
	return nil
}
