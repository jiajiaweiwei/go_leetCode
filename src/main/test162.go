package main

func main() {
	slice := []int{1, 2, 1, 3, 5, 6, 4}
	findPeakElement(slice)
}

// 二分查找 中等题 寻找峰值
// 方法1 使用二分查找 判断特定节点左右大小
func findPeakElement(nums []int) int {
	// 二分查找
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		//mid = (left + right) / 2 无法避免整数溢出
		mid = left + (right-left)/2 //可以避免整数溢出
		//判断节点是否符合条件
		//一直找最大的值
		if nums[mid] > nums[mid+1] { // 说明峰值在左侧或者mid
			right = mid
		} else {
			left = mid + 1 //说明峰值在右侧，
		}
	}
	return left
}

// 方法2 直接遍历 题目条件所给，表示一定有峰值，所以直接找最大值就是峰值之一
func findPeakElement2(nums []int) (resultNum int) {
	// 直接遍历
	// 返回值
	for i, v := range nums {
		if v > nums[resultNum] {
			resultNum = i
		}
	}
	return
}
