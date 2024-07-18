package main

import "fmt"

// 合并两个有序数组
func main() {
	nums1 := []int{0, 0, 0, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println("-------")
	merge(nums1, 0, nums2, 3)
	fmt.Println(nums1)
	fmt.Println(nums2)
}

// 88简单题 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 {
		copy(nums1, nums2)
		return
	}

	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 && k >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// 处理剩余元素
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
