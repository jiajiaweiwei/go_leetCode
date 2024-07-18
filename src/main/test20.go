package main

func main() {
	s := "([)]"
	isValid(s)
}

// 栈 简单题 有效的括号（判断括号是否有效）
// 方法1 栈
func isValid(s string) bool {
	// 先使用哈希表存起来
	hashMapLeft := make(map[rune]rune)
	hashMapRight := make(map[rune]rune)
	hashMapLeft['('] = ')'
	hashMapLeft['['] = ']'
	hashMapLeft['{'] = '}'
	hashMapRight[')'] = '('
	hashMapRight[']'] = '['
	hashMapRight['}'] = '{'
	// 使用栈判断
	stack := make([]rune, 0)
	for _, v := range s {
		// 是左括号 入栈
		if _, ok := hashMapLeft[v]; ok {
			stack = append(stack, v)
		} else { // 是右边括号 出栈一个元素并对比
			temp := hashMapRight[v]
			// 使用切片获取栈顶元素
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != temp {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
