package main

func main() {
	slice := []int{3, 4, 0, 1, 2}
	findMin(slice)
}

// 二分查找 中等题 寻找旋转排序数组中的最小值
// 方法1 二分查找 根据二分查找一半有序的特点 同理test162 寻找峰值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] < nums[right] { //最小值在左边
			right = mid
		} else { //最小值一定在有右边
			left = mid + 1
		}
	}
	return nums[left]
}
