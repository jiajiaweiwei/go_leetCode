package main

func main() {

}

// 力扣 第238中等题 除自身以外数组的乘积 （时间复杂度必须为On，不可用除法）
// 方法1 前缀之积  后缀之积 拆分时间复杂度
func productExceptSelf(nums []int) []int {
	length := len(nums)
	// L 和 R 分别表示左右两侧的乘积列表
	L, R, resultNum := make([]int, length), make([]int, length), make([]int, length)
	// L[i] 为索引 i 左侧所有元素的乘积
	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
	L[0] = 1
	R[length-1] = 1
	//求左侧乘积列表
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}
	//求右侧乘积列表
	for i := length - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	//求结果
	for i := 0; i < length; i++ {
		resultNum[i] = L[i] * R[i]
	}
	return resultNum
}

// 方法2 双指针 实现前缀积后缀积
func productExceptSelf2(nums []int) []int {
	length := len(nums)
	resultNum := make([]int, length)
	//初始化切片
	for i, _ := range resultNum {
		resultNum[i] = 1
	}
	//初始化前后缀之积
	beforeNum, afterNum := 1, 1
	for i, j := 0, length-1; i < length; i++ {
		resultNum[i] *= beforeNum
		resultNum[j] *= afterNum
		beforeNum *= nums[j]
		afterNum *= nums[j]
		j--
	}
	return resultNum

}
