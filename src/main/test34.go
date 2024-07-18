package main

func main() {

}

// 二分查找 中等题 找出排序数组中的第一个和最后一个位置
// 方法1 先找到一个位置，再根据有序性 遍历到最左边 再遍历到最右边 ，就是第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	// 特殊情况处理
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 返回值
	resultNum := make([]int, 2, 2)
	// 二分查找
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if target == nums[mid] {
			// 找到了其中一个位置
			// 开始判断哪个在前，哪个在后
			mid--
			// 先找到第一个位置
			for mid >= 0 && mid < len(nums) && target == nums[mid] {
				mid--
			}
			resultNum[0] = mid + 1
			mid++
			for mid >= 0 && mid < len(nums) && target == nums[mid] {
				mid++
			}
			resultNum[1] = mid - 1
			return resultNum
		}
		if target > nums[mid] {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}

// 方法2 直接两次二分查找，第一次第一个位置，第二次找最后一个位置 时间稍微比方法1 短
