package main

func main() {

}

// 力扣14 简单题 最长公共前缀
// 方法1 横向扫描 （暴力法） 每次遍历都更新最长公共前缀
func longestCommonPrefix1(strs []string) string {
	//特殊情况处理
	if len(strs) == 0 {
		return ""
	}
	preFix := strs[0]
	//遍历
	for i := 1; i < len(strs); i++ {
		preFix = getPreFix(preFix, strs[i])
	}
	return preFix

}

// 获取两个字符串的最长公共前缀
func getPreFix(str1, str2 string) string {
	var lenPre int
	if len(str1) > len(str2) {
		lenPre = len(str2)
	} else {
		lenPre = len(str1)
	}
	index := 0
	for index < lenPre && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

// 方法2 纵向扫描 （重点）（暴力法2） 遍历的是每一个字符串的某列
func longestCommonPrefix2(strs []string) string {
	//特殊情况处理
	if len(strs) == 0 {
		return ""
	}
	//纵向扫描
	//注意区分好内层循环和外层循环
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] { //strs[0][:i] 就表示对字符串strs[0]进行切片操作
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 方法3 分而治之
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 方法4 二分查找
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
