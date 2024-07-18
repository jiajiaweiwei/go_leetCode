package main

func main() {

}

// 多维动态规划 中等题 判断交错字符串
// 方法1 动态规划
func isInterleave(s1 string, s2 string, s3 string) bool {
	// 如果s1长度+s2长度 != s3长度 则一定不是交错字符串
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == len(s2) && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	// 动态规划变量
	f := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		f[i] = make([]bool, len(s2)+1)
	}
	f[0][0] = true
	// 状态转移方程 设置f[i][j]为s1前i个s2前j个能否组成s3前i+j个
	// f[i][j] = [i-1][j]&&s1[i-1]==s3[i + j - 1] || f[i][j-1]&&s2[j-1]==s3[i + j - 1]
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 {
				f[i][j] = f[i][j] || f[i-1][j] && (s1[i-1] == s3[i+j-1]) //or运算
			}
			if j > 0 {
				// ||运算能保存住f[i][j]的值，防止两次使用转台转移，导致后一次覆盖了前一次的值
				f[i][j] = f[i][j] || f[i][j-1] && (s2[j-1] == s3[i+j-1])
			}
		}
	}
	return f[len(s1)][len(s2)]
}

// 方法2 使用滚动数组优化空间复杂度
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	// 动态规划变量，使用一维数组代替二维数组
	f := make([]bool, len(s2)+1)
	f[0] = true
	// 初始化第一行
	for j := 1; j <= len(s2); j++ {
		f[j] = f[j-1] && s2[j-1] == s3[j-1] // 没判断s1时
	}
	for i := 1; i <= len(s1); i++ {
		// 更新第一列
		f[0] = f[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= len(s2); j++ {
			// 状态转移方程 加上了判断s1更新f[j]s1[i-1] == s3[i+j-1]
			f[j] = (f[j] && s1[i-1] == s3[i+j-1]) || (f[j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return f[len(s2)]
}
