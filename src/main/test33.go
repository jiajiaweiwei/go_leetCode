package main

func main() {

}

// 二分查找 中等题 搜索旋转排序数组
// 方法1  二分查找
// 原理 不管在哪里旋转的 将数组一分为2 总有一半是有序的
//
//	有序的部分使用二分 无序的继续一分为二 继续循环
func search33(nums []int, target int) int {
	// 将数组一分为二
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] > nums[right] { // 说明右边无序，左边有序
			// 左边有序，对左边二分
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else { // 目标值不在左边序列
				left = mid + 1
			}
		} else {
			// 右边有序，对右边二分
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else { // 目标值不在右边序列
				right = mid - 1
			}
		}
	}
	// 没找到到target就返回-1
	return -1
}
