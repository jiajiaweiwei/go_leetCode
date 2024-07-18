package main

func main() {
	wordDict := []string{"leet", "code"}
	wordBreak("leetcode", wordDict)
}

// 一维动态规划 中等题 单词拆分
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	// 动态规划变量
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			// 状态转移方程dp[i] = dp[j] && wordDictSet[s[j:i]]
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
