package main

// 力扣274 简单题 H指数（论文引用次数）
// 方法1 空间换时间
func hIndex(citations []int) int {
	//新创建一个切片
	l := len(citations)
	tempArry := make([]int, l+1)
	//将数组拷贝到切片中,并将元素值大于长度的置为长度
	for _, value := range citations {
		tempArry[min(value, l)]++
	}
	//0 1 2 3 4 5
	//3,0,6,1,5
	//1 1 0 1 0 2  tempArry结构
	//循环判断有哪些满足H指数的条件
	sum := 0
	for i := l; i >= 0; i-- { //从最大的开始判断
		sum += tempArry[i]
		if sum >= i {
			return i
		}
	}
	return 0
}

func min237(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 方法2 二分搜索法
func hIndex274(citations []int) int {
	//方法2 二分搜索法
	left, right := 0, len(citations)
	for left < right {
		mid := (left + right) + 1/2 //+1避免空数组导致死循环
		temp := 0
		for _, value := range citations {
			if value >= mid {
				temp++
			}
		}
		if temp >= mid {
			left = mid
		} else {
			right = mid - 1
		}

	}
	return left

}
func main() {
	testArray := []int{0}
	hIndex274(testArray)
}
