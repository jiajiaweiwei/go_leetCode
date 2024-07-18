package main

func main() {

}

// 数组简单题 判断特殊数组
// 方法1 一个个判断
func isArraySpecial(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	L := 1
	temp := nums[0] % 2
	for L < len(nums) {
		if nums[L]%2 == temp {
			return false
		}
		temp = nums[L] % 2
		L++
	}
	return true
}

// 方法2 直接两个两个判断
func isArraySpecial2(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1]%2 == nums[i]%2 {
			return false
		}
	}
	return true
}
