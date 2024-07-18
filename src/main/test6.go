package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	s2 := convert(s, 3)
	fmt.Println(s2)

}

// 力扣中等题Z型变换 （矩阵）
// 方法1 找到行索引的规律，按照规律将对应字符填入到新的切片中

func convert(s string, numRows int) string {
	//特殊情况处理
	if numRows < 2 || len(s) <= numRows {
		return s
	}
	l := len(s)
	resultString := make([][]byte, numRows)
	flag := 0
	down := true
	for i := 0; i < l; i++ {
		resultString[flag] = append(resultString[flag], s[i])
		//要使下标 012101210的循环
		if flag == 0 {
			down = true
		} else if flag == numRows-1 {
			down = false
		}
		if down {
			flag++
		} else {
			flag--
		}
	}
	temp := make([]string, numRows)
	fmt.Println(resultString)
	for i := 0; i < numRows; i++ {
		temp[i] = string(resultString[i])
	}
	return strings.Join(temp, "")
}
