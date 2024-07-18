package main

func main() {

}

// 二维矩阵查找 指定值
// 方法1 一次二分查找
func searchMatrix(matrix [][]int, target int) bool {
	//先计算出二维数组的总长度
	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1
	// 开始二分查找
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if target > matrix[mid/cols][mid%cols] {
			left = mid + 1
		} else if target < matrix[mid/cols][mid%cols] {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
