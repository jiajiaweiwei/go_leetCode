package main

func main() {
	s := "pwwkew"
	lengthOfLongestSubstring(s)
}

// 滑动窗口 中等题 无重复字符的最长子串

// 方法 滑动窗口+哈希
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针 并将元素放入map中 1表示出现，0表示未出现
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 自己实现的滑动窗口 重点 更加简单直接
func lengthOfLongestSubstring2(s string) int {
	hashMap := make(map[rune]int)
	//左指针
	left := 0
	//返回值
	max := 0
	//遍历字符串
	for right, char := range s {
		//如果发现哈希表中已经有了该字符，且在当前串中间val >= left 也就是在左指针右边 将原串直接跳过这个
		if val, ok := hashMap[char]; ok && val >= left {
			left = val + 1
		}
		//将右指针复制给哈希对应字符 值为字符的下标
		hashMap[char] = right
		//临时长度
		currentLength := right - left + 1
		if currentLength > max {
			max = currentLength
		}
	}
	return max
}
