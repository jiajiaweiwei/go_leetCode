package main

func main() {

}

// 矩阵 中等题 正方型矩阵旋转  要求原地旋转矩阵
// 重点方法  搞清楚元素的位置关系
// 方法 先水平翻转，后 主对角线翻转
func rotate48(matrix [][]int) {
	// 根据数学关系 目标元素 旋转后 a[i][j] -> a[j][i]
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
