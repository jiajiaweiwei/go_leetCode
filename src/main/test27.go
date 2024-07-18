package main

func main() {

}

// 力扣27 简单题 数组 移除元素
// 方法1 使用双指针
func removeElement(nums []int, val int) int {
	left := 0
	for _, num := range nums {
		if num != val {
			nums[left] = num
			left++
		}
	}
	return left
}

// 方法2 双指针优化 （优化了要移除的元素刚好在开头的可能性）
func removeElement1(nums []int, val int) int {
	// 左右指针向中间遍历
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}
