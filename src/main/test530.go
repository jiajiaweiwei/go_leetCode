package main

import (
	. "MyLikou/src/MyStruct"
	"math"
)

func main() {

}

// 二叉树 简单题 二叉搜索树的最小绝对差
// 方法1 使用中序遍历 二叉搜索树的中序遍历是由小到大
func getMinimumDifference(root *TreeNode) int {
	// 返回值
	resultNum, pre := math.MaxInt64, -1
	// 使用中序遍历匿名函数
	var midFS func(node *TreeNode)
	midFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		midFS(node.Left)
		if pre != -1 && node.Val-pre < resultNum {
			resultNum = node.Val - pre
		}
		pre = node.Val
		midFS(node.Right)
	}
	midFS(root)
	return resultNum
}
