package main

func main() {

}

// 哈希表 简单题存在重复元素2
// 方法1 使用哈希表
func containsNearbyDuplicate(nums []int, k int) bool {
	// 先存入哈希表
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := hashMap[nums[i]]; !ok {
			hashMap[nums[i]] = i
		} else if abs(i-v) <= k {
			return true
		} else {
			// 更新最近的下标
			hashMap[nums[i]] = i
		}
	}
	return false
}

// int型求绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// 方法2 滑动窗口  考虑数组 中的每个长度不超过 k+1k + 1k+1 的滑动窗口
func containsNearbyDuplicate1(nums []int, k int) bool {
	// 先存入哈希表
	hashMap := make(map[int]int)
	// 滑动窗口的思想
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(hashMap, nums[i-k-1])
		}
		if _, ok := hashMap[nums[i]]; ok {
			return true
		}
		hashMap[nums[i]] = i
	}
	return false
}
