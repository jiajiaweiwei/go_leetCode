package main

func main() {

}

// 力扣 189 中等题 轮转数组
// 方法1 使用额外的数组
func rotate(nums []int, k int) {
	temp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp[(k+i)%len(nums)] = nums[i]
	}
	copy(nums, temp)

}

// 方法2 环状替换
func rotate2(nums []int, k int) {

}

// 方法3 切片翻转 重点 好理解 原地翻转两次
func rotate3(nums []int, k int) {
	k = k % len(nums)
	reverse(nums[:])
	reverse(nums[:k])
	reverse(nums[k:])
}

// 重点 自定义翻转函数 将切片从i到j翻转
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 翻转函数的另外一种实现
func rotate4(nums []int, k int) {
	k = k % len(nums)
	reverse4(nums, 0, len(nums)-1)
	reverse4(nums, 0, k-1)
	reverse4(nums, k, len(nums)-1)
	return
}
func reverse4(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return
}
