package main

import "sort"

func main() {

}

// 哈希表 字母异位词 分组
// 方法1 对字符串排列后再判断哈希表中是否有对应的元素 将排完序的 作为键
func groupAnagrams(strs []string) [][]string {
	// 返回值
	resultString := make([][]string, 0)
	// 哈希表
	hashMap := make(map[string][]string)
	// 先将原字符串排序 存入哈希表中
	for i := 0; i < len(strs); i++ {
		strTemp := []byte(strs[i])
		sort.Slice(strTemp, func(i, j int) bool {
			return strTemp[i] < strTemp[j]
		})
		hashMap[string(strTemp)] = append(hashMap[string(strTemp)], strs[i])
	}
	// 将哈希表中数据存入返回值中
	for _, v := range hashMap {
		resultString = append(resultString, v)
	}
	return resultString
}

// 方法2 计数
