package main

import "sort"

func main() {

}

// 区间中等题 合并区间
// 方法1 排序
func merge56(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	for _, interval := range intervals {
		L, R := interval[0], interval[1]
		if len(merged) == 0 || merged[len(merged)-1][1] < L {
			merged = append(merged, []int{L, R})
		} else {
			merged[len(merged)-1][1] = max56(merged[len(merged)-1][1], R)
		}
	}

	return merged
}

func max56(x, y int) int {
	if x < y {
		return y
	}
	return x
}
