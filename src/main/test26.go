package main

func main() {

	nums := []int{1, 1}
	removeDuplicates(nums)

}

// 力扣26 简单题 删除数组中的重复项
// 方法1  普通快慢双指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//定义指针
	i, j := 0, 1
	for i < len(nums) && j < len(nums) {
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		if i+1 < len(nums) && j < len(nums) {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return i + 1
}

// 方法2 双指针优化
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
