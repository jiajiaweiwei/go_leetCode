package main

import "strconv"

func main() {

}

// 栈 中等题 逆波兰表达式
// 复习中缀表达式转后缀表达式（也就是逆波兰表达式）
// 从而得出逆波兰表达式 转 中缀表达式并计算的结果
// 方法1 使用栈
/*
数字的uncode表达范围
（） 的编码范围
+-* ·/的编码范围
*/
func evalRPN(tokens []string) int {
	// 栈
	stack := make([]int, 0)
	// 使用栈计算表达式
	for _, v := range tokens {
		num, err := strconv.Atoi(v) // 将字符串转化为整数
		// 是数字 入栈
		if err == nil {
			stack = append(stack, num)
		} else {
			// 不是数字 出栈两个元素 计算后入栈
			temp1 := stack[len(stack)-2]
			temp2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, temp1+temp2)
			case "-":
				stack = append(stack, temp1-temp2)
			case "*":
				stack = append(stack, temp1*temp2)
			default:
				stack = append(stack, temp1/temp2)
			}
		}
	}
	return stack[0]
}
