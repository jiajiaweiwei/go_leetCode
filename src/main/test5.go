package main

func main() {

}

// 多维动态规划 中等题 最长回文子串
// 方法1 两边向中间，动态规划 时间复杂度 On^2 空间复杂度 On^2
func longestPalindrome(s string) string {
	// 边界情况处理
	l := len(s)
	if l < 2 {
		return s
	}
	// 动态规划数组
	dp := make([][]bool, l)
	// 返回值数据
	resultNum := []rune(s)
	// 最大回文串的起点和终点
	start := 0
	maxLen := 1
	// 动态规划数组 初始化
	for i, _ := range dp {
		dp[i] = make([]bool, l)
	}
	// 初始化每一个元素都是回文串
	for i := 0; i < l; i++ {
		dp[i][i] = true
	}
	// 开始遍历 动态规划判断回文串
	// 串长度
	for L := 2; L <= l; L++ {
		// 串起点
		for i := 0; i < l; i++ {
			// 计算串终点
			j := L + i - 1
			// 如果是右边界就退出循环(字串判断完了)
			if j >= l {
				break
			}
			// 开始判断是否是回文串
			// 先判断首尾字符是否相同
			if resultNum[i] != resultNum[j] {
				// 不同，直接判断为false
				dp[i][j] = false
			} else {
				//相同就判断子串
				if j-i < 3 { // 串长度为3 时，可以不同判断子串
					dp[i][j] = true
				} else {
					//转台转移方程
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 更新最大长度的回文串
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	// 判断结束 根据起点终点输出返回结果
	return string(resultNum[start : maxLen+start])
}

// 方法2 中心向两边 中心扩展算法 时间复杂度 On^2 空间复杂度 O1
func longestPalindrome2(s string) string {
	// 返回值 (返回字符串的下标)
	//a,b int

	return "123"
}

// 方法3 Manacher 算法 时间复杂度 On  空间复杂度 O1
func longestPalindrome3(s string) string {
	// 返回值 (返回字符串的下标)
	//a,b int

	return "123"
}
