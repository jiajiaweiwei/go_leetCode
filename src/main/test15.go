package main

import (
	"fmt"
	"sort"
)

func main() {
	var tempNum = []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(tempNum)
	fmt.Println(tempNum)
	fmt.Println(threeSum(tempNum))

}

// 力扣中等题 三数之和 （有多解情况 且结果集不能重复）
// 方法1 先排序减低重复率 后遍历 根据有序性，将第二重循环和第三重循环可以将使用左右指针向中间遍历
func threeSum(nums []int) [][]int {
	resultNum := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// 避免重复枚举
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]
		j, k := i+1, len(nums)-1
		//优化，跳过不必要的步骤
		// 优化：如果当前元素与后面两个元素之和大于零，则结束循环
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// 优化：如果当前元素与数组中最后两个元素之和小于零，则跳过当前循环
		if nums[i]+nums[len(nums)-2]+nums[len(nums)-1] < 0 {
			continue
		}
		for j < k {
			sum := nums[j] + nums[k]
			if sum < target {
				j++
			} else if sum > target {
				k--
			} else {
				resultNum = append(resultNum, []int{nums[i], nums[j], nums[k]})
				// 避免重复
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过重复数字
			}
		}
	}
	return resultNum
}
