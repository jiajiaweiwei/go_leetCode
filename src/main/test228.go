package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 42
	s := strconv.Itoa(i)
	fmt.Printf("Integer %d converted to string: %s\n", i, s)
}

// 简单题 汇总区间
func summaryRanges(nums []int) []string {
	//返回值
	resultValue := make([]string, 0)
	//一次遍历
	for i := 0; i < len(nums); i++ {
		left := i
		for i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			i++
		}
		if i != left {
			resultValue = append(resultValue, strconv.Itoa(nums[left])+"->"+strconv.Itoa(nums[i]))
		} else {
			resultValue = append(resultValue, strconv.Itoa(nums[left]))
		}
	}
	return resultValue
}
