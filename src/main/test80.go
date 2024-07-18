package main

func main() {

}

// 力扣80  删除有序数组中的重复项 2 中等题
// 方法1 快慢双指针
func removeDuplicates88(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
