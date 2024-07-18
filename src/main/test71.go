package main

import "strings"

func main() {

}

// 栈 中等题  简化路径
// 方法1 使用栈
func simplifyPath(path string) string {
	var stack []string
	components := strings.Split(path, "/")

	for _, component := range components {
		if component == "" || component == "." {
			continue
		}
		if component == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, component)
		}
	}

	return "/" + strings.Join(stack, "/")
}
