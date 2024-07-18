package main

func main() {
	// 初始化一个整数切片
	slice := []int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}
	twoSum(slice, 542)
}

// 中等题 两数之和2 初始数组是严格地增的 空间复杂度要求为O1
func twoSum(numbers []int, target int) []int {
	//根据有序的特点使用二分查找
	//查找的值
	temp := 0
	//返回结果
	var resultNums = make([]int, 2, 2)
	for i := 0; i < len(numbers); i++ {
		for i < len(numbers)-1 && numbers[i] > target { //确定有唯一解，不需要考虑越界
			i++
		}
		temp = target - numbers[i]
		//使用二分查找
		l, r := 0, len(numbers)-1
		for l <= r {
			mid := (l + r) / 2
			if numbers[mid] == temp && mid != i {
				resultNums[0] = i + 1
				resultNums[1] = mid + 1
				return resultNums
			} else if mid == i {
				l++
			} else if numbers[l] > temp {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return resultNums
}
