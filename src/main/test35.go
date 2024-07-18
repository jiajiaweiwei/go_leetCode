package main

func main() {

}

// 二分查找 搜索插入位置
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	if nums[mid] < target {
		return left
	} else if nums[mid] > target {
		return mid - 1
	}
	return 0
}
