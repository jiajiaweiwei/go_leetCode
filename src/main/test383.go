package main

func main() {

}

// 哈希表 简单题 赎金信 判断字符构成
// 方法1 统计字符串中的字符频率，判断是否能够组成
func canConstruct(ransomNote string, magazine string) bool {
	// 先判断长度
	if len(ransomNote) > len(magazine) {
		return false
	}
	// 遍历 统计字符
	cnt := make([]int, 26, 26)
	for _, v := range magazine {
		cnt[v-'a']++
	}
	for _, v := range ransomNote {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
